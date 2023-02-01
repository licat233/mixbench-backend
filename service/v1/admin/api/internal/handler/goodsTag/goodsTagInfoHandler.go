package goodsTag

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mixbench/service/v1/admin/api/internal/logic/goodsTag"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
)

func GoodsTagInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GoodsTagGetReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := goodsTag.NewGoodsTagInfoLogic(r.Context(), svcCtx)
		resp, err := l.GoodsTagInfo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
