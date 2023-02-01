package goodsTag

import (
	"context"

	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodsTagPutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGoodsTagPutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodsTagPutLogic {
	return &GoodsTagPutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GoodsTagPutLogic) GoodsTagPut(req *types.GoodsTagPutReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
