package handler

import (
	"net/http"

	"mixbench/services/v1/admin/api/internal/logic"
	"mixbench/services/v1/admin/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func devPanelAuthHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewDevPanelAuthLogic(r.Context(), svcCtx, w, r)
		err := l.DevPanelAuth()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
