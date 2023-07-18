/*
 * @Author: licat
 * @Date: 2022-11-17 14:55:33
 * @LastEditors: licat
 * @LastEditTime: 2023-02-20 16:23:50
 * @Description: licat233@gmail.com
 */
package svc

import (
	"fmt"
	"log"
	"mixbench/common/mysqlx"
	"mixbench/services/v1/luckydraw/model"
	"mixbench/services/v1/luckydraw/mq/asynqmq"
	"mixbench/services/v1/luckydraw/mq/internal/config"
	luckydrawClient "mixbench/services/v1/luckydraw/rpc/client"

	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	LuckydrawRpc   *luckydrawClient.LuckydrawRpc
	UserModel      model.UserModel
	RedisClientOpt *asynq.RedisClientOpt
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn, err := mysqlx.NewMysql(c.Mysql.DataSource)
	if err != nil {
		log.Fatalln(fmt.Errorf("连接mysql数据库失败: %s", err))
	}
	var redisClientOpt = &asynq.RedisClientOpt{
		Addr:     c.Redis.Host,
		Password: c.Redis.Pass,
		DB:       asynqmq.RedisDB,
	}
	return &ServiceContext{
		Config:         c,
		LuckydrawRpc:   luckydrawClient.NewLuckydrawRpc(zrpc.MustNewClient(c.LuckydrawRpc)),
		UserModel:      model.NewUserModel(sqlConn, c.CacheRedis),
		RedisClientOpt: redisClientOpt,
	}
}
