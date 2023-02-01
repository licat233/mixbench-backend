package goodsTag

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mixbench/service/v1/admin/api/internal/logic/goodsTag"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
)

func GoodsTagListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GoodsTagListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := goodsTag.NewGoodsTagListLogic(r.Context(), svcCtx)
		resp, err := l.GoodsTagList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
