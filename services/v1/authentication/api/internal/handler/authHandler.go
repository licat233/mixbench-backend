package handler

import (
	"net/http"

	"mixbench/services/v1/authentication/api/internal/svc"
)

func authHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := svcCtx.AuthMiddlewareLogic.AuthHandler(w, r); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			w.WriteHeader(http.StatusOK)
		}
		// l := logic.NewAuthLogic(r.Context(), svcCtx, w, r)
		// err := l.Auth()
		// if err != nil {
		// 	w.WriteHeader(http.StatusUnauthorized)
		// 	// httpx.ErrorCtx(r.Context(), w, err);
		// } else {
		// 	w.WriteHeader(http.StatusOK)
		// 	// httpx.Ok(w)
		// }
	}
}
