package lucky

import (
	"context"

	"mixbench/services/v1/luckydraw/api/internal/svc"
	"mixbench/services/v1/luckydraw/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLuckyEnumsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLuckyEnumsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLuckyEnumsLogic {
	return &GetLuckyEnumsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLuckyEnumsLogic) GetLuckyEnums(req *types.GetLuckyEnumsReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
