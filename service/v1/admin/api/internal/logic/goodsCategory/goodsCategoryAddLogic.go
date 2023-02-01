package goodsCategory

import (
	"context"

	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodsCategoryAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGoodsCategoryAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodsCategoryAddLogic {
	return &GoodsCategoryAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GoodsCategoryAddLogic) GoodsCategoryAdd(req *types.GoodsCategoryAddReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
