/*
 * @Author: licat
 * @Date: 2023-02-17 15:41:34
 * @LastEditors: licat
 * @LastEditTime: 2023-02-20 16:04:44
 * @Description: licat233@gmail.com
 */
package svc

import (
	"log"
	"mixbench/common/rbacx"
	"mixbench/service/v1/admin/api/internal/config"
	"mixbench/service/v1/admin/api/internal/middleware"
	"mixbench/service/v1/admin/model"
	"mixbench/service/v1/admin/rpc/admin_pb"
	"mixbench/service/v1/captcha/rpc/captcha"
	"path"

	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
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

	CaptchaRpc     captcha.Captcha
	AdminRpc       AdminRpc
	CasbinEnforcer *casbin.SyncedEnforcer
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	db, err := sqlConn.RawDB()
	if err != nil {
		log.Fatalf("[NewServiceContext] sqlConn.RawDB err: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("[NewServiceContext] SqlConnPing err: %v", err)
	}
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
		Config:         c,
		AuthMiddleware: middleware.NewAuthMiddleware(c, JwtBlacklistsModel, AdminerModel).Handle,
		RbacMiddleware: middleware.NewRbacMiddleware(enforcer, &adminRpc.BaseClient).Handle,
		CaptchaRpc:     captcha.NewCaptcha(zrpc.MustNewClient(c.CaptchaRpc)),
		AdminRpc:       adminRpc,
		CasbinEnforcer: enforcer,
	}
}
