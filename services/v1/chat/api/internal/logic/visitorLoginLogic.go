package logic

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"mixbench/common/errorx"
	"mixbench/common/jwtx"
	"mixbench/common/respx"
	"mixbench/common/uniqueid"
	"mixbench/common/utils"
	"mixbench/services/v1/admin/rpc/admin_pb"
	"mixbench/services/v1/chat/api/internal/svc"
	"mixbench/services/v1/chat/api/internal/types"
	"mixbench/services/v1/chat/api/internal/ws"

	// "mixbench/services/v1/chat/api/internal/wshub"
	"mixbench/services/v1/chat/model"

	"github.com/asaskevich/govalidator"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mssola/useragent"
	"github.com/zeromicro/go-zero/core/logx"
)

type VisitorLoginLogic struct {
	logx.Logger
	ctx       context.Context
	svcCtx    *svc.ServiceContext
	r         *http.Request
	id        int64
	ip        string
	address   *utils.AddressInfo
	referer   string
	ua        *useragent.UserAgent
	language  string
	token     string
	ProjectId int64
}

func NewVisitorLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *VisitorLoginLogic {
	return &VisitorLoginLogic{
		Logger:    logx.WithContext(ctx),
		ctx:       ctx,
		svcCtx:    svcCtx,
		r:         r,
		id:        0,
		ip:        "",
		address:   nil,
		referer:   "",
		ua:        nil,
		language:  "",
		token:     "",
		ProjectId: 0,
	}
}

const VisitorTokenKey = "Authorization"

// 该接口主要用于生成 jwt-token
func (l *VisitorLoginLogic) VisitorLogin(req *types.VisitorLoginReq) (any, error) {
	token, _, err := l.Validator(strings.TrimSpace(l.r.Header.Get(VisitorTokenKey)))
	if err != nil {
		return nil, err
	}
	return respx.NewData(token), nil
}

func (l *VisitorLoginLogic) Validator(jwtToken string) (string, *model.Visitor, error) {
	// 获取设备信息
	l.ua = useragent.New(utils.RemoteUserAgent(l.r))
	//如果是机器人
	if l.ua.Bot() {
		return "", nil, errorx.ExternalError(errors.New("the current request is a bot"))
	}
	// 获取语言
	l.language = strings.Join(utils.RemoteLanguage(l.r), ",")

	// 获取 ip
	l.ip = utils.RemoteIp(l.r)
	// 解析 ip
	l.address = utils.QueryAddress(l.ip)

	// platform, err := utils.RemotePlatform(l.req)
	// if err != nil {
	// 	l.Logger.Error(err)
	// 	platform = fmt.Sprintf("%s[%s]", err.Error(), platform)
	// }
	// 获取来源
	weburl := l.r.Header.Get("Web-Url")
	if weburl != "" {
		l.referer = weburl
	}
	if l.referer == "" {
		l.referer = utils.RemoteReferer(l.r)
	}
	if l.referer == "" {
		l.referer = utils.RemoteOrigin(l.r)
	}
	if l.referer == "" {
		l.referer = "未知"
	}

	weburl = l.referer

	if weburl == "" || !govalidator.IsURL(weburl) {
		//未知referer
		return "", nil, errorx.AuthError("referer is " + weburl)
	}

	parsedURL, err := url.Parse(weburl)
	if err != nil {
		return "", nil, errorx.InternalError(fmt.Errorf("failed to parse URL: %s", err.Error()))
	}
	isDev := strings.HasPrefix(parsedURL.Host, "localhost")
	if !isDev {
		//获取当前weburl的归属项目
		webResp, err := l.svcCtx.AdminRpc.GetWebByUrl(l.ctx, &admin_pb.GetWebByUrlReq{
			Url: weburl,
		})
		if err != nil {
			return "", nil, err
		}

		if webResp.Web == nil {
			//如果无此url对应的web，则返回
			return "", nil, errorx.ExternalError("unknown referer")
		}
		l.ProjectId = webResp.Web.ProjectId
	} else {
		l.ProjectId = 1
	}

	// 获取游客可能存在的id
	var visitorId int64

	// 判断其是不是新客户，依据在于，是否携带jwt-token
	if jwtToken != "" {
		// 存在, 验证
		claims, ok := jwtx.ValidateToken(jwtToken, l.svcCtx.Config.VisitorAuthSecret)
		if ok {
			id, ok := jwtx.GetId(claims)
			if ok {
				//id值存在
				//合法性
				if id > 0 {
					//查一下这个id是否存在
					visitor, err := l.svcCtx.VisitorModel.FindById(l.ctx, id)
					if err != nil && err != model.ErrNotFound {
						l.Logger.Error(err)
						return "", nil, errorx.InternalError(err)
					}
					// 表示存在该visitor
					if err == nil {
						visitorId = id
						//查找项目是否存在
						projectResp, err := l.svcCtx.AdminRpc.GetProject(l.ctx, &admin_pb.GetProjectReq{Id: visitor.ProjectId})
						if err != nil {
							l.Logger.Error(err)
							return "", nil, errorx.InternalError(err)
						}
						if projectResp.Project == nil {
							//项目已不存在，游客将重新注册
							visitorId = -1
						} else {
							//如果存在，应该检验是否符合当前的project
							if l.ProjectId != visitor.ProjectId {
								//如果不符合，游客将重新注册
								visitorId = -1
							}
						}
					}
				}
			}
		}
		// err != nil 解析失败，该 token 不对，可能是恶意传的，不记录，算作新客户
	}

	l.id = visitorId

	// 新客户，给其分配客服:
	// ①轮询算法：按照客服的在线顺序，依次分配给游客。【不适合，会导致每个客服分配到的游客数量不一样】
	// ②随机算法：随机选择一个在线客服分配给游客。【不适合，有可能会出现每个客服分配的游客数量不一样】
	// ③最少负载算法：选择当前负载最小的客服分配给游客。【合适】
	// ④最近使用算法：选择最近使用过的客服分配给游客。【合适】

	// 最终方案：查看游客是否有历史客服id，如果存在，且该客服在线的话，则优先分配给该客服，否则，采用最小负载算法，如果最小负荷都一样，则采用随机算法。

	// 检查projectId
	if l.ProjectId < 1 {
		return "", nil, errorx.AuthError("project not found")
	}

	// 如果存在 visitorId，则表示老客户
	if visitorId > 0 {
		l.token = jwtToken
		fmt.Print("老用户")
		return l.isOldVisitor()
	} else {
		return l.isNewVisitor()
	}
}

func (l *VisitorLoginLogic) isOldVisitor() (string, *model.Visitor, error) {
	// 获取接待的客服id
	var kefuId int64
	// 获取其kefuId
	visitor, err := l.svcCtx.VisitorModel.FindById(l.ctx, l.id)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return "", nil, errorx.InternalError(err)
	}
	if err == model.ErrNotFound {
		// 表明不存在，将作为新客户处理
		return l.isNewVisitor()
	}
	//如果该游客已被拉黑 status == 2
	if visitor.Status != 1 && visitor.Status != 0 {
		return "", nil, errorx.ErrAuth
	}

	kefuId = visitor.KefuId

	//判断客服是否在线，如果不在线，给分配新的在线客服
	kefuOnline := ws.OnlineKefu.Exist(kefuId)
	if !kefuOnline {
		// 分配给其它客服
		newKefuId := ws.AssignKefu()
		// 表明有客服在线，则更新游客的接待客服，否则不更新，保留原来的客服
		if newKefuId > 0 && newKefuId != kefuId {
			//设置客服
			visitor.KefuId = newKefuId
		}
	}

	// 更新 visitor 信息
	visitor.LastLoginAt = time.Now().Local()
	if err := l.svcCtx.VisitorModel.Update(l.ctx, visitor); err != nil {
		l.Logger.Error(err)
		return "", nil, errorx.InternalError(err)
	}

	// 如果该客服在线，则没啥事了，因为这里的token是永久有效的，所以无需生产新的token
	return l.token, visitor, nil
}

func (l *VisitorLoginLogic) isNewVisitor() (string, *model.Visitor, error) {
	//表示为新客户，接下来为其分配客服
	kefuId := ws.AssignKefu()

	browser, _ := l.ua.Browser()

	name := "未知地區-網友"
	if l.address.Country != "" {
		name = l.address.Country
		if l.address.City != "" {
			name = fmt.Sprintf("%s-%s", name, l.address.City)
		}
		name = fmt.Sprintf("%s-網友", name)
	}

	// 构建 visitor 信息
	visitor := &model.Visitor{
		Id:          0,
		Uuid:        uniqueid.NewUUID(),
		Name:        name,
		Comment:     "",
		Ip:          l.ip,
		KefuId:      kefuId,
		ProjectId:   l.ProjectId,
		Status:      1,
		Referer:     l.referer,
		Region:      l.address.Country,
		City:        l.address.City,
		Platform:    l.ua.Platform(),
		Os:          l.ua.OS(),
		Browser:     browser,
		Language:    l.language,
		LastLoginAt: time.Now().Local(),
	}

	// 存入数据库
	res, err := l.svcCtx.VisitorModel.Insert(l.ctx, visitor)
	if err != nil {
		l.Logger.Error(err)
		return "", nil, errorx.InternalError(err)
	}

	//获取id
	newId, err := res.LastInsertId()
	if err != nil {
		l.Logger.Error(err)
		return "", nil, errorx.InternalError(err)
	}

	//取出数据库中的记录
	visitorMd, err := l.svcCtx.VisitorModel.FindById(l.ctx, newId)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return "", nil, errorx.InternalError(err)
	}
	if err == model.ErrNotFound {
		//见鬼了，可能是缓存未更新
		err = errors.New("the newly inserted data in the MySQL database is missing")
		l.Logger.Error(err)
		return "", nil, errorx.InternalError(err)
	}

	// 生成 token
	token, err := GenVisitorToken(l.svcCtx.Config.VisitorAuthSecret, visitorMd)
	if err != nil {
		l.Logger.Error(err)
		return "", nil, errorx.InternalError(err)
	}

	return token, visitorMd, nil
}

func GenVisitorToken(secret string, visitor *model.Visitor) (string, error) {
	claims := make(jwt.MapClaims)
	claims[jwtx.JwtUidKey] = visitor.Id
	claims[jwtx.JwtPlatformKey] = visitor.Platform
	return jwtx.CreateToken(claims, secret)
}
