package activity

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mixbench/services/v1/luckydraw/api/internal/logic/activity"
	"mixbench/services/v1/luckydraw/api/internal/svc"
)

func ActivityPatternEnumListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := activity.NewActivityPatternEnumListLogic(r.Context(), svcCtx)
		resp, err := l.ActivityPatternEnumList()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
