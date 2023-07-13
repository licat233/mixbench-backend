package middleware

import (
	adminClient "mixbench/services/v1/admin/rpc/client"
	"mixbench/services/v1/authentication/api/internal/config"
	"mixbench/services/v1/authentication/model"
	"net/http"
)

type AuthMiddleware struct {
}

func NewAuthMiddleware(c config.Config, jwtBlacklistsModel model.JwtBlacklistModel, adminRpc *adminClient.AdminRpc) *AuthMiddleware {
	return &AuthMiddleware{}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation

		// Passthrough to next handler if need
		next(w, r)
	}
}
