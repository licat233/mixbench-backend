package svc

import (
	"mixbench/services/v1/admin/api/internal/config"
	adminClient "mixbench/services/v1/admin/rpc/client"
	"mixbench/services/v1/authentication/rpc/authentication"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	AuthenticationRpc authentication.Authentication
	AdminRpc          *adminClient.AdminRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	adminRpc := adminClient.NewAdminRpc(zrpc.MustNewClient(c.AdminRpc))
	return &ServiceContext{
		Config:            c,
		AuthenticationRpc: authentication.NewAuthentication(zrpc.MustNewClient(c.AuthenticationRpc)),
		AdminRpc:          adminRpc,
	}
}
