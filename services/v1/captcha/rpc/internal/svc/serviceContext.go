package svc

import (
	"mixbench/services/v1/captcha/rpc/internal/config"
	"time"

	"github.com/mojocn/base64Captcha"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config        config.Config
	RedisClient   *redis.Redis
	CaptchaStore  base64Captcha.Store
	CaptchaExpire time.Duration
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisConn := redis.New(c.Redis.Host, func(r *redis.Redis) {
		r.Type = c.Redis.Type
		r.Pass = c.Redis.Pass
	})

	captchaExpire := time.Minute * 5
	return &ServiceContext{
		Config:        c,
		RedisClient:   redisConn,
		CaptchaStore:  base64Captcha.NewMemoryStore(10240, captchaExpire),
		CaptchaExpire: captchaExpire,
	}
}
