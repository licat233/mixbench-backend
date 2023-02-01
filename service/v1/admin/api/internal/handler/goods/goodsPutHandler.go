package goods

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mixbench/service/v1/admin/api/internal/logic/goods"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
)

func GoodsPutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GoodsPutReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := goods.NewGoodsPutLogic(r.Context(), svcCtx)
		resp, err := l.GoodsPut(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
