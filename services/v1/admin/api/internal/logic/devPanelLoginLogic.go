package logic

import (
	"context"
	"errors"
	"net/http"

	"mixbench/common/errorx"
	"mixbench/common/utils"
	"mixbench/services/v1/admin/api/internal/svc"
	"mixbench/services/v1/admin/api/internal/types"
	"mixbench/services/v1/admin/rpc/admin_pb"
	"mixbench/services/v1/captcha/rpc/captcha_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DevPanelLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	w      http.ResponseWriter
	r      *http.Request
}

func NewDevPanelLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext, w http.ResponseWriter, r *http.Request) *DevPanelLoginLogic {
	return &DevPanelLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		w:      w,
		r:      r,
	}
}

const cookieAuthName = "dev_token"

func (l *DevPanelLoginLogic) DevPanelLogin(req *types.DevPanelLoginReq) (resp *types.BaseResp, err error) {
	//获取请求设备
	reqPlatform, err := utils.RemotePlatform(l.r)
	if err != nil {
		return nil, errorx.ErrPlatform
	}
	//获取请求ip
	reqIp := utils.RemoteIp(l.r)

	//先进行验证码校验
	res, err := l.svcCtx.CaptchaRpc.VerifyCaptcha(l.ctx, &captcha_pb.VerifyCaptchaReq{
		CaptchaId: req.CaptchaId,
		Solution:  req.Solution,
	})
	if err != nil {
		return nil, err
	}
	if !res.Success {
		return nil, errorx.ExternalError(nil, errors.New("captcha auth failed"))
	}

	//验证码校验成功，调用admin.rpc服务登录
	jwtData := &admin_pb.JwtData{
		UserId:   0,
		Ip:       reqIp,
		Platform: reqPlatform,
	}
	respData, err := l.svcCtx.AdminRpc.SuperAdminLogin(l.ctx, &admin_pb.LoginReq{
		Passport:  req.Username,
		Password:  req.Password,
		AutoLogin: true,
		JwtData:   jwtData,
	})
	if err != nil {
		return nil, err
	}

	// 创建一个名为 "dev_authenticated" 的 Cookie
	cookie := &http.Cookie{
		Name:     cookieAuthName,
		Value:    respData.JwtToken.AccessToken,
		Domain:   ".mixbench.net", // 设置共享域的值（以点开头表示在 mixbench.net 及其子域中共享）
		Path:     "/",
		HttpOnly: false,
	}

	// 将 Cookie 设置到响应中
	http.SetCookie(l.w, cookie)

	// 同时设置 session
	session, _ := store.Get(l.r, sessionAuthKey)
	session.Values["authenticated"] = true
	session.Save(l.r, l.w)
	return
}
