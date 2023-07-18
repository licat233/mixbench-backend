package handler

import (
	"net/http"

	"mixbench/services/v1/authentication/api/internal/logic"
	"mixbench/services/v1/authentication/api/internal/svc"
)

func authHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewAuthLogic(r.Context(), svcCtx, w, r)
		err := l.Auth()
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			// httpx.ErrorCtx(r.Context(), w, err);
		} else {
			w.WriteHeader(http.StatusOK)
			// httpx.Ok(w)
		}
	}
}
