/*
 * @Author: licat
 * @Date: 2023-02-17 15:41:34
 * @LastEditors: licat
 * @LastEditTime: 2023-02-20 16:04:44
 * @Description: licat233@gmail.com
 */
package svc

import (
	"fmt"
	"mixbench/common/rbacx"
	"mixbench/common/utils"
	"mixbench/services/v1/admin/api/internal/config"
	"mixbench/services/v1/admin/api/internal/middleware"
	"mixbench/services/v1/admin/model"
	"mixbench/services/v1/admin/rpc/admin_pb"
	"mixbench/services/v1/authentication/rpc/authentication"
	"path"

	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type AdminRpc struct {
	admin_pb.AdminClient
	admin_pb.BaseClient
}

type ServiceContext struct {
	Config         config.Config
	AuthMiddleware rest.Middleware
	RbacMiddleware rest.Middleware

	AuthenticationRpc authentication.Authentication
	AdminRpc          AdminRpc
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

	clientConn := zrpc.MustNewClient(c.AdminRpc).Conn()

	adminRpc := AdminRpc{
		admin_pb.NewAdminClient(clientConn),
		admin_pb.NewBaseClient(clientConn),
	}
	casbinFile := path.Join(c.RootDir, "etc/rbac_with_domains_model.conf")
	enforcer := rbacx.NewCasbinEnforcer(db, casbinFile)
	return &ServiceContext{
		Config:            c,
		AuthMiddleware:    middleware.NewAuthMiddleware(c, JwtBlacklistsModel, AdminerModel).Handle,
		RbacMiddleware:    middleware.NewRbacMiddleware(enforcer, &adminRpc.BaseClient).Handle,
		AuthenticationRpc: authentication.NewAuthentication(zrpc.MustNewClient(c.AuthenticationRpc)),
		AdminRpc:          adminRpc,
		CasbinEnforcer:    enforcer,
		AdminerModel:      AdminerModel,
		JwtBlacklistModel: JwtBlacklistsModel,
	}
}
