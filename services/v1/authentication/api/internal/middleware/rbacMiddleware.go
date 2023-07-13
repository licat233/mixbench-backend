package middleware

import (
	"mixbench/services/v1/admin/rpc/admin_pb"
	"net/http"

	"github.com/casbin/casbin/v2"
)

type RbacMiddleware struct {
}

func NewRbacMiddleware(enforcer *casbin.SyncedEnforcer, adminBaseRpc *admin_pb.BaseClient) *RbacMiddleware {
	return &RbacMiddleware{}
}

func (m *RbacMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation

		// Passthrough to next handler if need
		next(w, r)
	}
}
