package middleware

import (
	"fmt"
	"mixbench/common/errorx"
	"mixbench/common/jwtx"
	"mixbench/services/v1/luckydraw/api/internal/config"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
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
		token := jwtx.GetToken(r)
		if token == "" {
			httpx.Error(w, errorx.ErrAuth)
			return
		}
		claims, ok := jwtx.ValidateToken(token, m.Config.VisitorAuthSecret)
		if !ok {
			httpx.Error(w, errorx.ErrAuth)
			return
		}

		id, exist := jwtx.GetId(claims)
		if !exist {
			httpx.Error(w, errorx.ErrAuth)
			return
		}
		r.Header.Add("uid", fmt.Sprint(id))
		next(w, r)
	}
}
