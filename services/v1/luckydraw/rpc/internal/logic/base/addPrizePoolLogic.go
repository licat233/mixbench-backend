package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/common/uniqueid"
	"mixbench/services/v1/luckydraw/model"
	"mixbench/services/v1/luckydraw/rpc/internal/svc"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPrizePoolLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddPrizePoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPrizePoolLogic {
	return &AddPrizePoolLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加活动奖品池，某活动中的奖品信息
func (l *AddPrizePoolLogic) AddPrizePool(in *luckydraw_pb.AddPrizePoolReq) (*luckydraw_pb.AddPrizePoolResp, error) {
	data := &model.PrizePool{
		Uuid:       uniqueid.NewUUID(),
		ActivityId: in.ActivityId,
		PrizeId:    in.PrizeId,
		GradeId:    in.GradeId,
		LuckValue:  in.LuckValue,
		Status:     in.Status,
		Total:      in.Total,
		Remainder:  in.Remainder,
	}
	_, err := l.svcCtx.PrizePoolModel.Insert(l.ctx, data)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &luckydraw_pb.AddPrizePoolResp{}, nil
}
