package svc

import (
	"mixbench/services/v1/admin/rpc/admin_pb"
	"mixbench/services/v1/authentication/api/internal/config"
	"mixbench/services/v1/authentication/rpc/authentication"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type AdminRpc struct {
	admin_pb.AdminClient
	admin_pb.BaseClient
}

type ServiceContext struct {
	Config            config.Config
	AuthenticationRpc authentication.Authentication
	AdminRpc          AdminRpc
	RedisClient       *redis.Redis
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
		Config:            c,
		AuthenticationRpc: authentication.NewAuthentication(zrpc.MustNewClient(c.AuthenticationRpc)),
		AdminRpc:          adminRpc,
		RedisClient:       redisConn,
	}
}
