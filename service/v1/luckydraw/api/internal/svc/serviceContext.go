package svc

import (
	"mixbench/service/v1/captcha/rpc/captcha"
	"mixbench/service/v1/luckydraw/api/internal/config"
	"mixbench/service/v1/luckydraw/api/internal/middleware"
	"mixbench/service/v1/luckydraw/rpc/luckydraw"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	CaptchaRpc    captcha.Captcha
	LuckdrawRpc   luckydraw.Luckdraw
	JwtMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		CaptchaRpc:    captcha.NewCaptcha(zrpc.MustNewClient(c.CaptchaRpc)),
		LuckdrawRpc:   luckydraw.NewLuckdraw(zrpc.MustNewClient(c.LuckdrawRpc)),
		JwtMiddleware: middleware.NewJwtMiddleware(c).Handle,
	}
}
