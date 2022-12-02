package prizeGrade

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mixbench/service/v1/admin/api/internal/logic/prizeGrade"
	"mixbench/service/v1/admin/api/internal/svc"
)

func PrizeGradeEnumListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := prizeGrade.NewPrizeGradeEnumListLogic(r.Context(), svcCtx)
		resp, err := l.PrizeGradeEnumList()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
