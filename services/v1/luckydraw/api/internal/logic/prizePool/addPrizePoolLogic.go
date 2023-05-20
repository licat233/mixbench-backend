package prizePool

import (
	"context"

	"mixbench/common/respx"
	"mixbench/services/v1/luckydraw/api/internal/svc"
	"mixbench/services/v1/luckydraw/api/internal/types"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPrizePoolLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddPrizePoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPrizePoolLogic {
	return &AddPrizePoolLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddPrizePoolLogic) AddPrizePool(req *types.AddPrizePoolReq) (any, error) {
	in := &luckydraw_pb.AddPrizePoolReq{
		ActivityId: req.ActivityId,
		PrizeId:    req.PrizeId,
		GradeId:    req.GradeId,
		LuckValue:  req.LuckValue,
		Status:     req.Status,
		Total:      req.Total,
		Remainder:  req.Remainder,
	}
	_, err := l.svcCtx.LuckydrawRpc.AddPrizePool(l.ctx, in)
	if err != nil {
		//若rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}

	return respx.DefaultStatusResp(nil)
}
