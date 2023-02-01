package goodsCategory

import (
	"context"

	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodsCategoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGoodsCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodsCategoryListLogic {
	return &GoodsCategoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GoodsCategoryListLogic) GoodsCategoryList(req *types.GoodsCategoryListReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
