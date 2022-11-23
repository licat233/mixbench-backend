package handler

import (
	"mixbench/service/v1/luckdraw/api/internal/logic"
	"mixbench/service/v1/luckdraw/api/internal/svc"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func luckViewHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewLuckViewLogic(r.Context(), svcCtx)
		err := l.LuckView()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
