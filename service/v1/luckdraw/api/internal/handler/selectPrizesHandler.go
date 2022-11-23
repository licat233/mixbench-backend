package handler

import (
	"mixbench/service/v1/luckdraw/api/internal/logic"
	"mixbench/service/v1/luckdraw/api/internal/svc"
	"mixbench/service/v1/luckdraw/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func selectPrizesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SelectPrizesReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewSelectPrizesLogic(r.Context(), svcCtx)
		resp, err := l.SelectPrizes(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
