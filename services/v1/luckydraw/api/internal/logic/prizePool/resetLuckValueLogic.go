package prizePool

import (
	"context"

	"mixbench/common/respx"
	"mixbench/services/v1/luckydraw/api/internal/svc"
	"mixbench/services/v1/luckydraw/api/internal/types"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResetLuckValueLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewResetLuckValueLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetLuckValueLogic {
	return &ResetLuckValueLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResetLuckValueLogic) ResetLuckValue(req *types.ResetLuckValueReq) (any, error) {
	_, err := l.svcCtx.LuckydrawRpc.ResetLuckValue(l.ctx, &luckydraw_pb.ResetLuckValueReq{
		ActivityId:   req.ActivityId,
		ActivityUuid: req.ActivityUuid,
	})
	if err != nil {
		return nil, err
	}
	return respx.StatusResp("", true), nil
}
