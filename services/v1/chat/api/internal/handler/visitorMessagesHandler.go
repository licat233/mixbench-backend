package handler

import (
	"net/http"

	"mixbench/services/v1/chat/api/internal/logic"
	"mixbench/services/v1/chat/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func VisitorMessagesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewVisitorMessagesLogic(r.Context(), svcCtx, r)
		resp, err := l.VisitorMessages()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
