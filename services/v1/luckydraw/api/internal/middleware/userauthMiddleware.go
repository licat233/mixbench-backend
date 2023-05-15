package middleware

import (
	"mixbench/services/v1/luckydraw/api/internal/config"
	"net/http"
)

type UserAuthMiddleware struct {
	Config config.Config
}

func NewUserAuthMiddleware(c config.Config) *UserAuthMiddleware {
	return &UserAuthMiddleware{
		Config: c,
	}
}

func (m *UserAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation

		// Passthrough to next handler if need
		next(w, r)
	}
}
