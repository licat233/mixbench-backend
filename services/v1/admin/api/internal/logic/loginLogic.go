/*
 * @Author: licat
 * @Date: 2023-02-17 15:41:34
 * @LastEditors: licat
 * @LastEditTime: 2023-02-20 12:23:04
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
	"mixbench/common/xcrypt"
	"mixbench/services/v1/admin/api/internal/svc"
	"mixbench/services/v1/admin/api/internal/types"
	"mixbench/services/v1/admin/rpc/admin_pb"
	"mixbench/services/v1/authentication/rpc/authentication_pb"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx     context.Context
	svcCtx  *svc.ServiceContext
	Request *http.Request
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *LoginLogic {
	return &LoginLogic{
		Logger:  logx.WithContext(ctx),
		ctx:     ctx,
		svcCtx:  svcCtx,
		Request: r,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (any, error) {
	//获取请求设备
	reqPlatform, err := utils.RemotePlatform(l.Request)
	if err != nil {
		return nil, errorx.ErrPlatform
	}
	//获取请求ip
	reqIp := utils.RemoteIp(l.Request)

	//调用captcha.rpc服务校验验证码
	res, err := l.svcCtx.AuthenticationRpc.VerifyCaptcha(l.ctx, &authentication_pb.VerifyCaptchaReq{
		CaptchaId: req.CaptchaId,
		Solution:  req.Solution,
	})
	if err != nil {
		return nil, err
	}
	if !res.Success {
		return nil, errorx.ExternalError(nil, errors.New("captcha auth failed"))
	}

	if err = xcrypt.PasswordValidator(req.Password); err != nil {
		return nil, errorx.ExtRpcErr(err)
	}

	//验证码校验成功，调用admin.rpc服务查询是否存在该账户
	var adminer *admin_pb.Adminer
	if govalidator.IsEmail(req.Passport) {
		resp, err := l.svcCtx.AdminRpc.FindByPassport(l.ctx, &admin_pb.FindByPassportReq{
			Passport: req.Passport,
		})
		if err != nil {
			return nil, err
		}
		adminer = resp.Adminer
	} else {
		resp, err := l.svcCtx.AdminRpc.FindByPassport(l.ctx, &admin_pb.FindByPassportReq{
			Passport: req.Passport,
		})
		if err != nil {
			return nil, err
		}
		adminer = resp.Adminer
	}

	if adminer == nil {
		return nil, errorx.ExternalError(nil, "account does not exist")
	}

	if !xcrypt.CheckPassword(req.Password, adminer.Password) {
		return nil, errorx.ExternalError(errors.New("account and password do not match"))
	}

	//检查adminer的状态
	if adminer.Status == 0 {
		return nil, errorx.ExternalError("this account has been disabled")
	}

	//生成token
	jwtData := &authentication_pb.JwtData{
		UserId:   adminer.Id,
		Ip:       reqIp,
		Platform: reqPlatform,
	}
	resp, err := l.svcCtx.AuthenticationRpc.GenToken(l.ctx, &authentication_pb.GenTokenReq{
		JwtData:   jwtData,
		AutoLogin: req.AutoLogin,
	})
	if err != nil {
		return nil, err
	}

	if resp.JwtToken == nil {
		return nil, errorx.InternalError("JwtToken equal nil")
	}

	//更新登录数据
	adminer.LastLoginAt = time.Now().Local().Unix()
	adminer.LoginCount++
	_, err = l.svcCtx.AdminRpc.PutAdminer(l.ctx, &admin_pb.PutAdminerReq{
		Adminer: adminer,
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
