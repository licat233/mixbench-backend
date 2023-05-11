package handler

import (
	"net/http"

	"mixbench/service/v1/admin/api/internal/logic"
	"mixbench/service/v1/admin/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func authHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewAuthLogic(r.Context(), svcCtx, r)
		err := l.Auth()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
