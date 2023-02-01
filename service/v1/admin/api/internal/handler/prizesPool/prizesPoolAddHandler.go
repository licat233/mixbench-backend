package prizesPool

import (
	"net/http"

	"mixbench/service/v1/admin/api/internal/logic/prizesPool"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func PrizePoolAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PrizePoolAddReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := prizesPool.NewPrizePoolAddLogic(r.Context(), svcCtx)
		resp, err := l.PrizePoolAdd(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
