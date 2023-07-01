package handler

import (
	"net/http"

	"mixbench/services/v1/chat/api/internal/logic"
	"mixbench/services/v1/chat/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func VisitorServerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewVisitorServerLogic(r.Context(), svcCtx, w, r)
		err := l.VisitorServer()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
