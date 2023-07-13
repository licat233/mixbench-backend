package web

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mixbench/services/v1/admin/api/internal/logic/web"
	"mixbench/services/v1/admin/api/internal/svc"
	"mixbench/services/v1/admin/api/internal/types"
)

func PutWebHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PutWebReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := web.NewPutWebLogic(r.Context(), svcCtx)
		resp, err := l.PutWeb(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
