package lucky

import (
	"net/http"

	"mixbench/service/v1/admin/api/internal/logic/lucky"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func LuckyDetailsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LuckyDetailsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := lucky.NewLuckyDetailsLogic(r.Context(), svcCtx)
		resp, err := l.LuckyDetails(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
