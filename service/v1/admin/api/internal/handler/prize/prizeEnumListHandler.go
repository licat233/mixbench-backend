package prize

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mixbench/service/v1/admin/api/internal/logic/prize"
	"mixbench/service/v1/admin/api/internal/svc"
)

func PrizeEnumListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := prize.NewPrizeEnumListLogic(r.Context(), svcCtx)
		resp, err := l.PrizeEnumList()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
