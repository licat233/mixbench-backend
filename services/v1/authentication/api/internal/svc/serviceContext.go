package svc

import (
	"fmt"
	"log"
	"mixbench/common/rbacx"
	"mixbench/common/utils"
	adminClient "mixbench/services/v1/admin/rpc/client"
	"mixbench/services/v1/authentication/api/internal/config"
	"mixbench/services/v1/authentication/api/internal/middleware"
	"mixbench/services/v1/authentication/model"
	authenticationClient "mixbench/services/v1/authentication/rpc/client"
	"path"

	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	AuthMiddleware rest.Middleware
	RbacMiddleware rest.Middleware

	AuthenticationRpc *authenticationClient.AuthenticationRpc
	AdminRpc          *adminClient.AdminRpc
	RedisClient       *redis.Redis

	JwtBlacklistModel model.JwtBlacklistModel

	CasbinEnforcer *casbin.SyncedEnforcer
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn, err := utils.NewMysql(c.Mysql.DataSource)
	if err != nil {
		log.Fatalln(fmt.Errorf("连接mysql数据库失败: %s", err))
	}
	redisConn := redis.New(c.Redis.Host, func(r *redis.Redis) {
		r.Type = c.Redis.Type
		r.Pass = c.Redis.Pass
	})
	// 因为 utils.NewMysql 方法内部已经执行过 sqlConn.RawDB() 方法且无错误了，这里再执行选择忽略错误
	db, _ := sqlConn.RawDB()
	JwtBlacklistsModel := model.NewJwtBlacklistModel(sqlConn, c.CacheRedis)
	casbinFile := path.Join(c.RootDir, "etc/rbac_with_domains_model.conf")
	enforcer := rbacx.NewCasbinEnforcer(db, casbinFile)
	adminRpc := adminClient.NewAdminRpc(zrpc.MustNewClient(c.AuthenticationRpc))
	return &ServiceContext{
		Config:            c,
		AuthMiddleware:    middleware.NewAuthMiddleware(c, JwtBlacklistsModel, adminRpc).Handle,
		RbacMiddleware:    middleware.NewRbacMiddleware(enforcer, &adminRpc.BaseClient).Handle,
		AuthenticationRpc: authenticationClient.NewAuthenticationRpc(zrpc.MustNewClient(c.AuthenticationRpc)),
		AdminRpc:          adminRpc,
		RedisClient:       redisConn,
		JwtBlacklistModel: JwtBlacklistsModel,
		CasbinEnforcer:    enforcer,
	}
}
