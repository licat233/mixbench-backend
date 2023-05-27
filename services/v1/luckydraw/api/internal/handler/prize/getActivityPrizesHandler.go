package prize

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mixbench/services/v1/luckydraw/api/internal/logic/prize"
	"mixbench/services/v1/luckydraw/api/internal/svc"
	"mixbench/services/v1/luckydraw/api/internal/types"
)

func GetActivityPrizesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetActivityPrizesReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := prize.NewGetActivityPrizesLogic(r.Context(), svcCtx)
		resp, err := l.GetActivityPrizes(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
