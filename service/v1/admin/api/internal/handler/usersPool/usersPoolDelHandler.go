package usersPool

import (
	"net/http"

	"mixbench/service/v1/admin/api/internal/logic/usersPool"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserPoolDelHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserPoolDelReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := usersPool.NewUserPoolDelLogic(r.Context(), svcCtx)
		resp, err := l.UserPoolDel(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
