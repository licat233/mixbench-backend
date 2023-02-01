package goods

import (
	"context"

	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodsPutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGoodsPutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodsPutLogic {
	return &GoodsPutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GoodsPutLogic) GoodsPut(req *types.GoodsPutReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
