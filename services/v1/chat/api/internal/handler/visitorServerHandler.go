package handler

import (
	"net/http"

	"mixbench/common/errorx"
	"mixbench/common/utils"
	"mixbench/services/v1/chat/api/internal/logic"
	"mixbench/services/v1/chat/api/internal/svc"
	"mixbench/services/v1/chat/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func VisitorServerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !utils.IsWebsocketRequest(r) {
			err := errorx.ExternalError("the current request is not a WebSocket request")
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		var req types.ConnWebsocketReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := logic.NewVisitorServerLogic(r.Context(), svcCtx, w, r)
		_ = l.VisitorServer(&req)
		// if err != nil {
		// 	httpx.ErrorCtx(r.Context(), w, err)
		// }
		// if err == nil {
		// 	httpx.Ok(w)
		// }

		// 注意：在 WebSocket 通信结束后，不要再对 ResponseWriter 进行操作

	}
}
