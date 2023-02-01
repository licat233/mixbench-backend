package goodsCategory

import (
	"context"

	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodsCategoryDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGoodsCategoryDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodsCategoryDelLogic {
	return &GoodsCategoryDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GoodsCategoryDelLogic) GoodsCategoryDel(req *types.GoodsCategoryDelReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
