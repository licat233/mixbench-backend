/*
 * @Author: licat
 * @Date: 2022-11-17 14:55:33
 * @LastEditors: licat
 * @LastEditTime: 2023-02-20 16:23:50
 * @Description: licat233@gmail.com
 */
package svc

import (
	"mixbench/service/v1/luckydraw"
	"mixbench/service/v1/luckydraw/asynqmq"
	"mixbench/service/v1/luckydraw/model"
	"mixbench/service/v1/luckydraw/mq/internal/config"

	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	LuckydrawRpc   luckydraw.LuckydrawRpc
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
		LuckydrawRpc:   luckydraw.NewLuckydrawRpc(zrpc.MustNewClient(c.LuckydrawRpc)),
		UserModel:      model.NewUserModel(sqlConn, c.CacheRedis),
		RedisClientOpt: redisClientOpt,
	}
}
