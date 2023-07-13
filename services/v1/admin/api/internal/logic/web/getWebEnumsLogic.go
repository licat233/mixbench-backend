package web

import (
	"context"

	"mixbench/services/v1/admin/api/internal/svc"
	"mixbench/services/v1/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetWebEnumsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetWebEnumsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWebEnumsLogic {
	return &GetWebEnumsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetWebEnumsLogic) GetWebEnums(req *types.GetWebEnumsReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
