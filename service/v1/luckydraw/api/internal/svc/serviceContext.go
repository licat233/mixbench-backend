/*
 * @Author: licat
 * @Date: 2022-11-17 14:55:33
 * @LastEditors: licat
 * @LastEditTime: 2023-02-20 18:32:45
 * @Description: licat233@gmail.com
 */
package svc

import (
	"mixbench/service/v1/captcha/rpc/captcha"
	"mixbench/service/v1/luckydraw"
	"mixbench/service/v1/luckydraw/api/internal/config"
	"mixbench/service/v1/luckydraw/api/internal/middleware"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config              config.Config
	CaptchaRpc          captcha.Captcha
	LuckydrawRpc        luckydraw.LuckydrawRpc
	AdminAuthMiddleware rest.Middleware
	UserAuthMiddleware  rest.Middleware
	RbacMiddleware      rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:              c,
		CaptchaRpc:          captcha.NewCaptcha(zrpc.MustNewClient(c.CaptchaRpc)),
		LuckydrawRpc:        luckydraw.NewLuckydrawRpc(zrpc.MustNewClient(c.LuckydrawRpc)),
		AdminAuthMiddleware: middleware.NewAdminAuthMiddleware(c).Handle,
		UserAuthMiddleware:  middleware.NewUserAuthMiddleware(c).Handle,
		RbacMiddleware:      nil,
	}
}
