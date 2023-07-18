package svc

import (
	"context"
	"fmt"
	"log"
	"mixbench/common/mysqlx"
	"mixbench/services/v1/authentication/model"
	"mixbench/services/v1/authentication/mq/asynqmq"
	"mixbench/services/v1/authentication/rpc/internal/config"
	"time"

	"github.com/hibiken/asynq"
	"github.com/mojocn/base64Captcha"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config        config.Config
	RedisClient   *redis.Redis
	CaptchaStore  base64Captcha.Store
	CaptchaExpire time.Duration

	KqPusherClient *kq.Pusher
	AsynqClient    *asynqmq.AsynqProducer

	TokenBlacklistModel model.TokenBlacklistModel
	IpBlacklistModel    model.IpBlacklistModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn, err := mysqlx.NewMysql(c.Mysql.DataSource)
	if err != nil {
		log.Fatalln(fmt.Errorf("连接mysql数据库失败: %s", err))
	}
	redisConn := redis.New(c.Redis.Host, func(r *redis.Redis) {
		r.Type = c.Redis.Type
		r.Pass = c.Redis.Pass
	})

	//验证码有效期
	captchaExpire := time.Duration(c.Captcha.Expire*1000) * time.Millisecond
	return &ServiceContext{
		Config:              c,
		RedisClient:         redisConn,
		CaptchaStore:        base64Captcha.NewMemoryStore(10240, captchaExpire),
		CaptchaExpire:       captchaExpire,
		KqPusherClient:      &kq.Pusher{},
		AsynqClient:         asynqmq.NewAsynqProducer(context.Background(), &asynq.RedisClientOpt{Addr: c.Redis.Host, Password: c.Redis.Pass, DB: asynqmq.RedisDB}),
		TokenBlacklistModel: model.NewTokenBlacklistModel(sqlConn, c.CacheRedis),
		IpBlacklistModel:    model.NewIpBlacklistModel(sqlConn, c.CacheRedis),
	}
}
