package handler

import (
	"mixbench/service/v1/luckydraw/api/internal/logic"
	"mixbench/service/v1/luckydraw/api/internal/svc"
	"mixbench/service/v1/luckydraw/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func luckyListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LuckyListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewLuckyListLogic(r.Context(), svcCtx)
		resp, err := l.LuckyList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
