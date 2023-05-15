package handler

import (
	"net/http"

	"mixbench/services/v1/admin/api/internal/logic"
	"mixbench/services/v1/admin/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func authHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewAuthLogic(r.Context(), svcCtx, w, r)
		err := l.Auth()
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			// httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
