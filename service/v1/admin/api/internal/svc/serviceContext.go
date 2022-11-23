package svc

import (
	"mixbench/service/v1/admin/api/internal/config"
	"mixbench/service/v1/admin/api/internal/middleware"
	"mixbench/service/v1/admin/rpc/admin"
	"mixbench/service/v1/captcha/rpc/captcha"
	"mixbench/service/v1/model"

	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	JwtMiddleware  rest.Middleware
	RbacMiddleware rest.Middleware

	CaptchaRpc captcha.Captcha
	AdminRpc   admin.Admin

	CasbinEnforcer *casbin.SyncedEnforcer
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	db, _ := sqlConn.RawDB()
	rbacMiddleware := middleware.NewRbacMiddleware(db, c)
	JwtBlacklistsModel := model.NewJwtBlacklistsModel(sqlConn, c.CacheRedis)
	AdminerModel := model.NewAdminerModel(sqlConn, c.CacheRedis)
	return &ServiceContext{
		Config:         c,
		JwtMiddleware:  middleware.NewJwtMiddleware(c, JwtBlacklistsModel, AdminerModel).Handle,
		RbacMiddleware: rbacMiddleware.Handle,
		CaptchaRpc:     captcha.NewCaptcha(zrpc.MustNewClient(c.CaptchaRpc)),
		AdminRpc:       admin.NewAdmin(zrpc.MustNewClient(c.AdminRpc)),
		CasbinEnforcer: rbacMiddleware.CasbinEnforcer,
	}
}
