package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/services/v1/luckydraw/rpc/internal/svc"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelPrizePoolLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelPrizePoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelPrizePoolLogic {
	return &DelPrizePoolLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除活动奖品池，某活动中的奖品信息
func (l *DelPrizePoolLogic) DelPrizePool(in *luckydraw_pb.DelPrizePoolReq) (*luckydraw_pb.DelPrizePoolResp, error) {
	err := l.svcCtx.PrizePoolModel.Delete(l.ctx, in.Id)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &luckydraw_pb.DelPrizePoolResp{}, nil
}
