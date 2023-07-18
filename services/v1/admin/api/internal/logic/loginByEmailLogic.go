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
	"time"

	"mixbench/common/errorx"
	"mixbench/common/respx"
	"mixbench/common/utils"
	"mixbench/services/v1/admin/api/internal/svc"
	"mixbench/services/v1/admin/api/internal/types"
	"mixbench/services/v1/admin/rpc/admin_pb"
	"mixbench/services/v1/authentication/rpc/authentication_pb"

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
	res, err := l.svcCtx.AuthenticationRpc.VerifyEmailCaptcha(l.ctx, &authentication_pb.VerifyEmailCaptchaReq{
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
	//验证码校验成功，调用admin.rpc服务查询是否存在该email绑定的账户
	resp, err := l.svcCtx.AdminRpc.FindByEmail(l.ctx, &admin_pb.FindByEmailReq{
		Email: req.Email,
	})
	if err != nil {
		return nil, err
	}
	if resp.Adminer == nil {
		return nil, errorx.ExternalError(nil, "account does not exist")
	}

	//如果用户存在，则生成token
	jwtData := &authentication_pb.JwtData{
		UserId:   resp.Adminer.Id,
		Ip:       reqIp,
		Platform: reqPlatform,
	}
	resp1, err := l.svcCtx.AuthenticationRpc.GenToken(l.ctx, &authentication_pb.GenTokenReq{
		JwtData:   jwtData,
		AutoLogin: req.AutoLogin,
	})
	if err != nil {
		return nil, err
	}

	if resp1.JwtToken == nil {
		return nil, errorx.InternalError("JwtToken equal nil")
	}

	//更新登录数据
	adminer := resp.Adminer
	adminer.LastLoginAt = time.Now().Local().Unix()
	adminer.LoginCount++
	_, err = l.svcCtx.AdminRpc.PutAdminer(l.ctx, &admin_pb.PutAdminerReq{
		Adminer: adminer,
	})
	if err != nil {
		return nil, err
	}

	return respx.NewData(&types.JwtToken{
		AccessToken:  resp1.JwtToken.AccessToken,
		AccessExpire: resp1.JwtToken.AccessExpire,
		RefreshAfter: resp1.JwtToken.RefreshAfter,
	}), nil
}
