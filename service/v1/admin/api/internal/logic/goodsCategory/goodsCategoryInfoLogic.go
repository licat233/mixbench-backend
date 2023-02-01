package goodsCategory

import (
	"context"

	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodsCategoryInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGoodsCategoryInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodsCategoryInfoLogic {
	return &GoodsCategoryInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GoodsCategoryInfoLogic) GoodsCategoryInfo(req *types.GoodsCategoryGetReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
