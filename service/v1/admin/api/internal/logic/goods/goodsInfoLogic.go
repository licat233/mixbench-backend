package goods

import (
	"context"

	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodsInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGoodsInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodsInfoLogic {
	return &GoodsInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GoodsInfoLogic) GoodsInfo(req *types.GoodsGetReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
