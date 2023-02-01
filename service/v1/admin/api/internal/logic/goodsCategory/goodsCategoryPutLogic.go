package goodsCategory

import (
	"context"

	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodsCategoryPutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGoodsCategoryPutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodsCategoryPutLogic {
	return &GoodsCategoryPutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GoodsCategoryPutLogic) GoodsCategoryPut(req *types.GoodsCategoryPutReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
