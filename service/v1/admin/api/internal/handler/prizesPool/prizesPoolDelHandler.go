package prizesPool

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mixbench/service/v1/admin/api/internal/logic/prizesPool"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
)

func PrizePoolDelHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PrizePoolDelReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := prizesPool.NewPrizePoolDelLogic(r.Context(), svcCtx)
		resp, err := l.PrizePoolDel(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
