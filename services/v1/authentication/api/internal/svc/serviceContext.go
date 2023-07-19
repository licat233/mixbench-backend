package svc

import (
	"fmt"
	"log"
	"mixbench/common/mysqlx"
	"mixbench/common/rbacx"
	adminClient "mixbench/services/v1/admin/rpc/client"
	"mixbench/services/v1/authentication/api/internal/config"
	"mixbench/services/v1/authentication/api/internal/middleware"
	"mixbench/services/v1/authentication/rpc/authentication"

	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	AuthMiddleware rest.Middleware
	RbacMiddleware rest.Middleware

	AuthMiddlewareLogic *middleware.AuthMiddleware

	AuthenticationRpc authentication.Authentication
	AdminRpc          *adminClient.AdminRpc
	RedisClient       *redis.Redis

	CasbinEnforcer *casbin.SyncedEnforcer
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn, err := mysqlx.NewMysql(c.Mysql.DataSource)
	if err != nil {
		log.Fatalln(fmt.Errorf("连接mysql数据库失败: %s", err))
	}
	redisConn := redis.New(c.Redis.Host, func(r *redis.Redis) {
		r.Type = c.Redis.Type
		r.Pass = c.Redis.Pass
	})
	// 因为 mysqlx.NewMysql 方法内部已经执行过 sqlConn.RawDB() 方法且无错误了，这里再执行选择忽略错误
	db, _ := sqlConn.RawDB()

	enforcer := rbacx.NewCasbinEnforcer(db)
	adminRpc := adminClient.NewAdminRpc(zrpc.MustNewClient(c.AdminRpc))

	authMiddlewareLogic := middleware.NewAuthMiddleware(c, redisConn, adminRpc)
	return &ServiceContext{
		Config:              c,
		AuthMiddleware:      authMiddlewareLogic.Handle,
		RbacMiddleware:      middleware.NewRbacMiddleware(enforcer, &adminRpc.BaseClient).Handle,
		AuthMiddlewareLogic: authMiddlewareLogic,
		AuthenticationRpc:   authentication.NewAuthentication(zrpc.MustNewClient(c.AuthenticationRpc)),
		AdminRpc:            adminRpc,
		RedisClient:         redisConn,
		CasbinEnforcer:      enforcer,
	}
}
