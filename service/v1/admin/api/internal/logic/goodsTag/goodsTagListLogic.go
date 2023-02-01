package goodsTag

import (
	"context"

	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodsTagListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGoodsTagListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodsTagListLogic {
	return &GoodsTagListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GoodsTagListLogic) GoodsTagList(req *types.GoodsTagListReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
