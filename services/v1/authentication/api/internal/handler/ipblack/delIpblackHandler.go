package ipblack

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mixbench/services/v1/authentication/api/internal/logic/ipblack"
	"mixbench/services/v1/authentication/api/internal/svc"
	"mixbench/services/v1/authentication/api/internal/types"
)

func DelIpblackHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DelIpblackReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := ipblack.NewDelIpblackLogic(r.Context(), svcCtx)
		resp, err := l.DelIpblack(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
