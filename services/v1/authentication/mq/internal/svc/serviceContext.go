/*
 * @Author: licat
 * @Date: 2022-11-17 14:55:33
 * @LastEditors: licat
 * @LastEditTime: 2023-02-20 16:23:50
 * @Description: licat233@gmail.com
 */
package svc

import (
	"mixbench/services/v1/authentication/mq/asynqmq"
	"mixbench/services/v1/authentication/mq/internal/config"
	"mixbench/services/v1/authentication/rpc/authentication"

	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config            config.Config
	AuthenticationRpc authentication.Authentication
	RedisClientOpt    *asynq.RedisClientOpt
	RedisClient       *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	var redisClientOpt = &asynq.RedisClientOpt{
		Addr:     c.Redis.Host,
		Password: c.Redis.Pass,
		DB:       asynqmq.RedisDB,
	}
	redisCli := redis.New(c.Redis.Host, func(r *redis.Redis) {
		r.Type = c.Redis.Type
		r.Pass = c.Redis.Pass
	})
	return &ServiceContext{
		Config:            c,
		AuthenticationRpc: authentication.NewAuthentication(zrpc.MustNewClient(c.AuthenticationRpc)),
		RedisClientOpt:    redisClientOpt,
		RedisClient:       redisCli,
	}
}
