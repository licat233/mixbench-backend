package handler

import (
	"mixbench/service/v1/luckdraw/api/internal/logic"
	"mixbench/service/v1/luckdraw/api/internal/svc"
	"mixbench/service/v1/luckdraw/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func prizesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PrizesReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewPrizesLogic(r.Context(), svcCtx)
		resp, err := l.Prizes(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
