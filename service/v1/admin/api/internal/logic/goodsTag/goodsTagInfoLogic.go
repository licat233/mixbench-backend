package goodsTag

import (
	"context"

	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodsTagInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGoodsTagInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodsTagInfoLogic {
	return &GoodsTagInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GoodsTagInfoLogic) GoodsTagInfo(req *types.GoodsTagGetReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
