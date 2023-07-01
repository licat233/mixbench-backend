package handler

import (
	"net/http"

	"mixbench/services/v1/chat/api/internal/logic"
	"mixbench/services/v1/chat/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func KefuServerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewKefuServerLogic(r.Context(), svcCtx, w, r)
		err := l.KefuServer()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
