package logic

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"mixbench/common/errorx"
	"mixbench/common/jwtx"
	"mixbench/common/respx"
	"mixbench/common/uniqueid"
	"mixbench/common/utils"
	"mixbench/services/v1/chat/api/internal/svc"
	"mixbench/services/v1/chat/api/internal/types"
	"mixbench/services/v1/chat/api/internal/wshub"
	"mixbench/services/v1/chat/model"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mssola/useragent"
	"github.com/zeromicro/go-zero/core/logx"
)

type VisitorLoginLogic struct {
	logx.Logger
	ctx      context.Context
	svcCtx   *svc.ServiceContext
	req      *http.Request
	id       int64
	ip       string
	address  *utils.AddressInfo
	referer  string
	ua       *useragent.UserAgent
	language string
}

func NewVisitorLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext, req *http.Request) *VisitorLoginLogic {
	return &VisitorLoginLogic{
		Logger:   logx.WithContext(ctx),
		ctx:      ctx,
		svcCtx:   svcCtx,
		req:      req,
		id:       0,
		ip:       "",
		address:  nil,
		referer:  "",
		ua:       nil,
		language: "",
	}
}

const VisitorTokenKey = "jwt-token"
const KefuJwtKey = "kefuId"

// 该接口主要用于生成 jwt-token
func (l *VisitorLoginLogic) VisitorLogin(req *types.VisitorLoginReq) (any, error) {
	// 获取设备信息
	l.ua = useragent.New(utils.RemoteUserAgent(l.req))
	//如果是机器人
	if l.ua.Bot() {
		return nil, errorx.ExternalError(errors.New("the current request is a bot"))
	}
	// 获取语言
	l.language = strings.Join(utils.RemoteLanguage(l.req), ",")

	// 获取 ip
	l.ip = utils.RemoteIp(l.req)
	// 解析 ip
	l.address = utils.QueryAddress(l.ip)

	// platform, err := utils.RemotePlatform(l.req)
	// if err != nil {
	// 	l.Logger.Error(err)
	// 	platform = fmt.Sprintf("%s[%s]", err.Error(), platform)
	// }
	// 获取来源
	l.referer = utils.RemoteReferer(l.req)

	// 获取游客可能存在的id
	var visitorId int64

	// 判断其是不是新客户，依据在于，是否携带jwt-token
	jwtToken := l.req.Header.Get(VisitorTokenKey)
	jwtToken = strings.TrimSpace(jwtToken)
	if jwtToken != "" {
		// 存在，解析
		claims, err := jwtx.ParseToken(jwtToken, l.svcCtx.Config.VisitorAuthSecret)
		if err == nil {
			visitorIdV := jwtx.GetClaimsValue("id", *claims)
			if visitorIdV != "" {
				id, err := strconv.ParseInt(visitorIdV, 10, 64)
				if err != nil {
					// 值错误，算作新客户
					l.Logger.Error(err)
				} else {
					// visitorId = id
					//查一下这个id是否存在
					_, err = l.svcCtx.VisitorModel.FindOne(l.ctx, id)
					if err != nil && err != model.ErrNotFound {
						l.Logger.Error(err)
						return nil, errorx.InternalError(err)
					}
					//表示存在该visitor
					if err == nil {
						visitorId = id
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

	// 如果存在 visitorId，则表示老客户
	if visitorId > 0 {
		return l.isOldVisitor()
	} else {
		return l.isNewVisitor()
	}
}

func (l *VisitorLoginLogic) isOldVisitor() (any, error) {
	// 获取接待的客服id
	var kefuId int64
	// 获取其kefuId
	visitor, err := l.svcCtx.VisitorModel.FindOne(l.ctx, l.id)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.InternalError(err)
	}
	if err == model.ErrNotFound {
		// 表明不存在，将作为新客户处理
		return l.isNewVisitor()
	}
	kefuId = visitor.KefuId

	kefuCli := wshub.HubSvc.KefuClients.Get(kefuId)
	//客服如果不在线，
	if kefuCli == nil {
		// 分配给其它客服
		kefuCli = wshub.HubSvc.AssignKefu()
		// 表明没有客服在线
		if kefuCli == nil {
			// 则保留其原来的客服, 不更新visitor
		} else {
			// 更新 visitor 信息
			visitor.KefuId = kefuId
			if err := l.svcCtx.VisitorModel.Update(l.ctx, visitor); err != nil {
				l.Logger.Error(err)
				return nil, errorx.InternalError(err)
			}
			//此时需要更新聊天记录，不然新客服无法看到过往的聊天记录
			transferLogic := NewTransferVisitorLogic(l.ctx, l.svcCtx)
			_, err := transferLogic.TransferVisitor(&types.TransferVisitorReq{
				VisitorId: 0,
				KefuId:    kefuId,
			})
			if err != nil {
				return nil, err
			}
		}
	}

	// 如果该客服在线，则没啥事了，直接生成新的 token 返回
	token, err := GenVisitorToken(l.svcCtx.Config.VisitorAuthSecret, visitor)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.InternalError(err)
	}
	return respx.NewData(token), nil
}

func (l *VisitorLoginLogic) isNewVisitor() (any, error) {
	// 获取接待的客服id
	var kefuId int64
	//表示为新客户，接下来为其分配客服
	kefuCli := wshub.HubSvc.AssignKefu()
	// 表明没有客服在线
	if kefuCli == nil {
		// 没有客服在线，将其id设为 -1 表示其没有分配到客服
		kefuId = -1
	} else {
		kefuId = kefuCli.Kefu.Id
	}

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
		Id:       kefuId,
		Uuid:     uniqueid.NewUUID(),
		Name:     name,
		Ip:       l.ip,
		KefuId:   kefuId,
		Status:   1,
		Referer:  l.referer,
		Region:   l.address.Country,
		City:     l.address.City,
		Platform: l.ua.Platform(),
		Os:       l.ua.OS(),
		Browser:  browser,
		Language: l.language,
	}

	// 生成 token
	token, err := GenVisitorToken(l.svcCtx.Config.VisitorAuthSecret, visitor)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.InternalError(err)
	}

	// 存入数据库
	_, err = l.svcCtx.VisitorModel.Insert(l.ctx, visitor)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.InternalError(err)
	}

	return respx.NewData(token), nil
}

func GenVisitorToken(secret string, visitor *model.Visitor) (string, error) {
	claims := make(jwt.MapClaims)
	claims[KefuJwtKey] = visitor.KefuId
	claims["id"] = visitor.Id
	claims["platform"] = visitor.Platform
	return jwtx.CreateToken(claims, secret)
}
