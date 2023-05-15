/*
 * @Author: licat
 * @Date: 2023-02-18 16:47:41
 * @LastEditors: licat
 * @LastEditTime: 2023-02-20 12:37:10
 * @Description: licat233@gmail.com
 */
package middleware

import (
	"context"
	"mixbench/services/v1/admin/rpc/admin_pb"
	"net/http"

	"github.com/casbin/casbin/v2"
)

type RbacMiddleware struct {
	ctx            context.Context
	CasbinEnforcer *casbin.SyncedEnforcer
	AdminBaseRpc   *admin_pb.BaseClient
}

func NewRbacMiddleware(enforcer *casbin.SyncedEnforcer, adminBaseRpc *admin_pb.BaseClient) *RbacMiddleware {
	return &RbacMiddleware{
		ctx:            context.Background(),
		CasbinEnforcer: enforcer,
		AdminBaseRpc:   adminBaseRpc,
	}
}

func (m *RbacMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation

		// Passthrough to next handler if need
		next(w, r)
	}
}
