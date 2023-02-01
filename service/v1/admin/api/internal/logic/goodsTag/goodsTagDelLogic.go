package goodsTag

import (
	"context"

	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodsTagDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGoodsTagDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodsTagDelLogic {
	return &GoodsTagDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GoodsTagDelLogic) GoodsTagDel(req *types.GoodsTagDelReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
