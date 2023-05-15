package prizePool

import (
	"net/http"

	"mixbench/services/v1/luckydraw/api/internal/logic/prizePool"
	"mixbench/services/v1/luckydraw/api/internal/svc"
	"mixbench/services/v1/luckydraw/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetPrizePoolListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetPrizePoolListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := prizePool.NewGetPrizePoolListLogic(r.Context(), svcCtx)
		resp, err := l.GetPrizePoolList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
