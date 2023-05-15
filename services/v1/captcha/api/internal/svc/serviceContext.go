/*
 * @Author: licat
 * @Date: 2022-11-17 14:55:33
 * @LastEditors: licat
 * @LastEditTime: 2023-02-20 16:08:52
 * @Description: licat233@gmail.com
 */
package svc

import (
	"mixbench/services/v1/admin/rpc/admin_pb"
	"mixbench/services/v1/captcha/api/internal/config"
	"mixbench/services/v1/captcha/rpc/captcha"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type AdminRpc struct {
	admin_pb.AdminClient
	admin_pb.BaseClient
}

type ServiceContext struct {
	Config      config.Config
	CaptchaRpc  captcha.Captcha
	AdminRpc    AdminRpc
	RedisClient *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisConn := redis.New(c.Redis.Addr, func(r *redis.Redis) {
		r.Type = c.Redis.Type
		r.Pass = c.Redis.Pass
		r.Addr = c.Redis.Addr
	})

	clientConn := zrpc.MustNewClient(c.AdminRpc).Conn()

	adminRpc := AdminRpc{
		admin_pb.NewAdminClient(clientConn),
		admin_pb.NewBaseClient(clientConn),
	}

	return &ServiceContext{
		Config:      c,
		CaptchaRpc:  captcha.NewCaptcha(zrpc.MustNewClient(c.CaptchaRpc)),
		AdminRpc:    adminRpc,
		RedisClient: redisConn,
	}
}
