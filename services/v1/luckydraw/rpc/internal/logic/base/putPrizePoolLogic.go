package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/services/v1/luckydraw/model"
	"mixbench/services/v1/luckydraw/rpc/internal/svc"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutPrizePoolLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPutPrizePoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutPrizePoolLogic {
	return &PutPrizePoolLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新活动奖品池，某活动中的奖品信息
func (l *PutPrizePoolLogic) PutPrizePool(in *luckydraw_pb.PutPrizePoolReq) (*luckydraw_pb.PutPrizePoolResp, error) {
	data := &model.PrizePool{
		Id:         in.Id,
		Uuid:       in.Uuid,
		ActivityId: in.ActivityId,
		PrizeId:    in.PrizeId,
		GradeId:    in.GradeId,
		LuckValue:  in.LuckValue,
		Status:     in.Status,
		Total:      in.Total,
		Remainder:  in.Remainder,
	}

	if err := l.svcCtx.PrizePoolModel.Update(l.ctx, data); err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &luckydraw_pb.PutPrizePoolResp{}, nil
}
