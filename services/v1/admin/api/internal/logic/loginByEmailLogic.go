/*
 * @Author: licat
 * @Date: 2023-02-17 15:41:34
 * @LastEditors: licat
 * @LastEditTime: 2023-02-20 11:21:43
 * @Description: licat233@gmail.com
 */
package logic

import (
	"context"
	"errors"
	"net/http"

	"mixbench/common/errorx"
	"mixbench/common/respx"
	"mixbench/common/utils"
	"mixbench/services/v1/admin/api/internal/svc"
	"mixbench/services/v1/admin/api/internal/types"
	"mixbench/services/v1/admin/rpc/admin_pb"
	"mixbench/services/v1/captcha/rpc/captcha_pb"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginByEmailLogic struct {
	logx.Logger
	ctx     context.Context
	svcCtx  *svc.ServiceContext
	Request *http.Request
}

func NewLoginByEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *LoginByEmailLogic {
	return &LoginByEmailLogic{
		Logger:  logx.WithContext(ctx),
		ctx:     ctx,
		svcCtx:  svcCtx,
		Request: r,
	}
}

func (l *LoginByEmailLogic) LoginByEmail(req *types.LoginByEmailReq) (any, error) {
	//邮箱格式验证
	if !govalidator.IsEmail(req.Email) {
		return nil, errorx.ErrInvalidEmail
	}

	//获取请求设备
	reqPlatform, err := utils.RemotePlatform(l.Request)
	if err != nil {
		return nil, errorx.ErrPlatform
	}
	//获取请求ip
	reqIp := utils.RemoteIp(l.Request)

	//调用captcha.rpc服务校验验证码
	res, err := l.svcCtx.CaptchaRpc.AuthEmailCaptcha(l.ctx, &captcha_pb.AuthEmailCaptchaReq{
		Email:     req.Email,
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
	resp, err := l.svcCtx.AdminRpc.LoginByEmail(l.ctx, &admin_pb.LoginByEmailReq{
		Email:     req.Email,
		AutoLogin: req.AutoLogin,
		JwtData:   jwtData,
	})

	if err != nil {
		return nil, err
	}

	return respx.NewData(&types.JwtToken{
		AccessToken:  resp.JwtToken.AccessToken,
		AccessExpire: resp.JwtToken.AccessExpire,
		RefreshAfter: resp.JwtToken.RefreshAfter,
	}), nil
}
