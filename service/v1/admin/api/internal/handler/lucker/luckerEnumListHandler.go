package lucker

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mixbench/service/v1/admin/api/internal/logic/lucker"
	"mixbench/service/v1/admin/api/internal/svc"
)

func LuckerEnumListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := lucker.NewLuckerEnumListLogic(r.Context(), svcCtx)
		resp, err := l.LuckerEnumList()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
