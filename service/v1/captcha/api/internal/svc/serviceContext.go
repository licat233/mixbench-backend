package svc

import (
	"mixbench/service/v1/captcha/api/internal/config"
	"mixbench/service/v1/captcha/rpc/captcha"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	CaptchaRpc  captcha.Captcha
	RedisClient *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisConn := redis.New(c.Redis.Addr, func(r *redis.Redis) {
		r.Type = c.Redis.Type
		r.Pass = c.Redis.Pass
		r.Addr = c.Redis.Addr
	})

	return &ServiceContext{
		Config:      c,
		CaptchaRpc:  captcha.NewCaptcha(zrpc.MustNewClient(c.CaptchaRpc)),
		RedisClient: redisConn,
	}
}
