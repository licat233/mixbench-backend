package ipblack

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mixbench/services/v1/authentication/api/internal/logic/ipblack"
	"mixbench/services/v1/authentication/api/internal/svc"
	"mixbench/services/v1/authentication/api/internal/types"
)

func GetIpblackHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetIpblackReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := ipblack.NewGetIpblackLogic(r.Context(), svcCtx)
		resp, err := l.GetIpblack(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
