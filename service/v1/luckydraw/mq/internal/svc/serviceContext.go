/*
 * @Author: licat
 * @Date: 2022-11-17 14:55:33
 * @LastEditors: licat
 * @LastEditTime: 2023-01-30 17:51:43
 * @Description: licat233@gmail.com
 */
package svc

import (
	"mixbench/common/asynqmq"
	"mixbench/service/v1/luckydraw/mq/internal/config"
	"mixbench/service/v1/luckydraw/rpc/luckydraw"
	"mixbench/service/v1/model"

	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	LuckdrawRpc    luckydraw.Luckdraw
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
		LuckdrawRpc:    luckydraw.NewLuckdraw(zrpc.MustNewClient(c.LuckdrawRpc)),
		UserModel:      model.NewUserModel(sqlConn, c.CacheRedis),
		RedisClientOpt: redisClientOpt,
	}
}
