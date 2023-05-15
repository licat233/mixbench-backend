package prizePool

import (
	"context"

	"mixbench/services/v1/luckydraw/api/internal/svc"
	"mixbench/services/v1/luckydraw/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPrizePoolEnumsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPrizePoolEnumsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPrizePoolEnumsLogic {
	return &GetPrizePoolEnumsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPrizePoolEnumsLogic) GetPrizePoolEnums(req *types.GetPrizePoolEnumsReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
