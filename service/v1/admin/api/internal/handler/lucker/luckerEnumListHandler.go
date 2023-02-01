package lucky

import (
	"net/http"

	"mixbench/service/v1/admin/api/internal/logic/lucky"
	"mixbench/service/v1/admin/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func LuckyEnumListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := lucky.NewLuckyEnumListLogic(r.Context(), svcCtx)
		resp, err := l.LuckyEnumList()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
