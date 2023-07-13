package svc

import (
	"fmt"
	"mixbench/common/rbacx"
	"mixbench/common/utils"
	"mixbench/services/v1/admin/api/internal/config"
	"mixbench/services/v1/admin/api/internal/middleware"
	"mixbench/services/v1/admin/model"
	adminClient "mixbench/services/v1/admin/rpc/client"
	authenticationClient "mixbench/services/v1/authentication/rpc/client"
	"path"

	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	AuthMiddleware rest.Middleware
	RbacMiddleware rest.Middleware

	AuthenticationRpc *authenticationClient.AuthenticationRpc
	AdminRpc          *adminClient.AdminRpc
	CasbinEnforcer    *casbin.SyncedEnforcer

	AdminerModel      model.AdminerModel
	JwtBlacklistModel model.JwtBlacklistModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn, err := utils.NewMysql(c.Mysql.DataSource)
	if err != nil {
		panic(fmt.Errorf("连接mysql数据库失败: %s", err))
	}
	// 因为 utils.NewMysql 方法内部已经执行过 sqlConn.RawDB() 方法且无错误了，这里再执行选择忽略错误
	db, _ := sqlConn.RawDB()
	JwtBlacklistsModel := model.NewJwtBlacklistModel(sqlConn, c.CacheRedis)
	AdminerModel := model.NewAdminerModel(sqlConn, c.CacheRedis)

	casbinFile := path.Join(c.RootDir, "etc/rbac_with_domains_model.conf")
	enforcer := rbacx.NewCasbinEnforcer(db, casbinFile)
	adminRpc := adminClient.NewAdminRpc(zrpc.MustNewClient(c.AdminRpc))
	return &ServiceContext{
		Config:            c,
		AuthMiddleware:    middleware.NewAuthMiddleware(c, JwtBlacklistsModel, AdminerModel).Handle,
		RbacMiddleware:    middleware.NewRbacMiddleware(enforcer, &adminRpc.BaseClient).Handle,
		AuthenticationRpc: authenticationClient.NewAuthenticationRpc(zrpc.MustNewClient(c.AuthenticationRpc)),
		AdminRpc:          adminRpc,
		CasbinEnforcer:    enforcer,
		AdminerModel:      AdminerModel,
		JwtBlacklistModel: JwtBlacklistsModel,
	}
}
