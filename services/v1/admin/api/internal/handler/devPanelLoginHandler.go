package handler

import (
	"net/http"

	"mixbench/services/v1/admin/api/internal/logic"
	"mixbench/services/v1/admin/api/internal/svc"
	"mixbench/services/v1/admin/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func devPanelLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DevPanelLoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewDevPanelLoginLogic(r.Context(), svcCtx, w, r)
		resp, err := l.DevPanelLogin(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
