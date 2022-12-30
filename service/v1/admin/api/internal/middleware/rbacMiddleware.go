/*
 * @Author: licat
 * @Date: 2022-11-17 14:55:32
 * @LastEditors: licat
 * @LastEditTime: 2022-12-30 11:54:47
 * @Description: licat233@gmail.com
 */
package middleware

import (
	"context"
	"database/sql"
	"mixbench/service/v1/admin/api/internal/config"
	"mixbench/service/v1/admin/rpc/admin"
	"net/http"
	"path"

	sqladapter "github.com/Blank-Xu/sql-adapter"
	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
)

type RbacMiddleware struct {
	ctx            context.Context
	CasbinEnforcer *casbin.SyncedEnforcer
	AdminRpc       admin.Admin
}

func NewRbacMiddleware(db *sql.DB, c config.Config) *RbacMiddleware {
	casbinFile := path.Join(c.RootDir, "etc/rbac_with_domains_model.conf")
	ctx := context.Background()
	return &RbacMiddleware{
		ctx:            context.Background(),
		CasbinEnforcer: NewCasbinEnforcer(ctx, db, casbinFile),
		AdminRpc:       admin.NewAdmin(zrpc.MustNewClient(c.AdminRpc)),
	}
}

func (m *RbacMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 1.获取用户id
		// 2.查询用户权限

		// Passthrough to next handler if need
		next(w, r)
	}
}

func NewCasbinEnforcer(ctx context.Context, db *sql.DB, casbinFile string) *casbin.SyncedEnforcer {
	if db == nil {
		return nil
	}
	adapter, err := sqladapter.NewAdapter(db, "mysql", "casbin_rule")
	if err != nil {
		logx.WithContext(ctx).Errorf("新建Casbin Adapter失败:%v", err)
		return nil
	}

	syncedEnforcer, err := casbin.NewSyncedEnforcer(casbinFile, adapter)
	if err != nil {
		logx.WithContext(ctx).Errorf("新建Casbin Enforcer失败:%v", err)
		return nil
	}

	// Load the policy from DB.
	if err = syncedEnforcer.LoadPolicy(); err != nil {
		logx.WithContext(ctx).Errorf("加载Casbin Policy失败:%v", err)
		return nil
	}

	return syncedEnforcer
}
