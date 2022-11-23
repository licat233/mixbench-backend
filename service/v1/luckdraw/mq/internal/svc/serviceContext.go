package svc

import (
	"mixbench/common/asynqmq"
	"mixbench/service/v1/luckdraw/mq/internal/config"
	"mixbench/service/v1/luckdraw/rpc/luckdraw"
	"mixbench/service/v1/model"

	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	LuckdrawRpc    luckdraw.Luckdraw
	UserModel      model.UserModel
	RedisClientOpt *asynq.RedisClientOpt
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	var redisClientOpt = &asynq.RedisClientOpt{
		Addr:     c.Redis.Host,
		Password: c.Redis.Pass,
		DB:       asynqmq.RedisDB,
	}
	return &ServiceContext{
		Config:         c,
		LuckdrawRpc:    luckdraw.NewLuckdraw(zrpc.MustNewClient(c.LuckdrawRpc)),
		UserModel:      model.NewUserModel(sqlConn, c.CacheRedis),
		RedisClientOpt: redisClientOpt,
	}
}
