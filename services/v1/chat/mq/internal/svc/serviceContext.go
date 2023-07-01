/*
 * @Author: licat
 * @Date: 2022-11-17 14:55:33
 * @LastEditors: licat
 * @LastEditTime: 2023-02-20 16:23:50
 * @Description: licat233@gmail.com
 */
package svc

import (
	"mixbench/services/v1/chat/model"
	"mixbench/services/v1/chat/mq/asynqmq"
	"mixbench/services/v1/chat/mq/internal/config"

	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config         config.Config
	MessageModel   model.MessageModel
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
		MessageModel:   model.NewMessageModel(sqlConn, c.CacheRedis),
		RedisClientOpt: redisClientOpt,
	}
}
