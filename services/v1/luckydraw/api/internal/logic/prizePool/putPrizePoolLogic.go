package prizePool

import (
	"context"

	"mixbench/common/respx"
	"mixbench/services/v1/luckydraw/api/internal/svc"
	"mixbench/services/v1/luckydraw/api/internal/types"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutPrizePoolLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPutPrizePoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutPrizePoolLogic {
	return &PutPrizePoolLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutPrizePoolLogic) PutPrizePool(req *types.PutPrizePoolReq) (any, error) {
	in := &luckydraw_pb.PutPrizePoolReq{
		Id:         req.Id,
		Uuid:       req.Uuid,
		ActivityId: req.ActivityId,
		PrizeId:    req.PrizeId,
		GradeId:    req.GradeId,
		LuckValue:  req.LuckValue,
		Status:     req.Status,
		Total:      req.Total,
		Remainder:  req.Remainder,
	}
	_, err := l.svcCtx.LuckydrawRpc.PutPrizePool(l.ctx, in)
	if err != nil {
		//若rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}

	return respx.DefaultStateResp(nil)
}
