package handler

import (
	"net/http"

	"mixbench/services/v1/admin/api/internal/logic"
	"mixbench/services/v1/admin/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func currentUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewCurrentUserLogic(r.Context(), svcCtx, r)
		resp, err := l.CurrentUser()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
