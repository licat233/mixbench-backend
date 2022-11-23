package activity

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mixbench/service/v1/admin/api/internal/logic/activity"
	"mixbench/service/v1/admin/api/internal/svc"
)

func ActivityPatternEnumListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := activity.NewActivityPatternEnumListLogic(r.Context(), svcCtx)
		resp, err := l.ActivityPatternEnumList()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
