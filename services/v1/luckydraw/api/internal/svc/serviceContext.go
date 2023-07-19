/*
 * @Author: licat
 * @Date: 2022-11-17 14:55:33
 * @LastEditors: licat
 * @LastEditTime: 2023-02-20 18:32:45
 * @Description: licat233@gmail.com
 */
package svc

import (
	"mixbench/services/v1/authentication/rpc/authentication"
	"mixbench/services/v1/luckydraw/api/internal/config"
	"mixbench/services/v1/luckydraw/api/internal/middleware"
	luckydrawClient "mixbench/services/v1/luckydraw/rpc/client"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	AuthenticationRpc authentication.Authentication
	LuckydrawRpc      *luckydrawClient.LuckydrawRpc

	UserAuthMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:             c,
		AuthenticationRpc:  authentication.NewAuthentication(zrpc.MustNewClient(c.AuthenticationRpc)),
		LuckydrawRpc:       luckydrawClient.NewLuckydrawRpc(zrpc.MustNewClient(c.LuckydrawRpc)),
		UserAuthMiddleware: middleware.NewUserAuthMiddleware(c).Handle,
	}
}
