package userPool

import (
	"context"

	"mixbench/services/v1/luckydraw/api/internal/svc"
	"mixbench/services/v1/luckydraw/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserPoolEnumsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserPoolEnumsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserPoolEnumsLogic {
	return &GetUserPoolEnumsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserPoolEnumsLogic) GetUserPoolEnums(req *types.GetUserPoolEnumsReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
