package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/service/v1/luckydraw/model"
	"mixbench/service/v1/luckydraw/rpc/internal/logic/dataconv"
	"mixbench/service/v1/luckydraw/rpc/internal/svc"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPrizePoolLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPrizePoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPrizePoolLogic {
	return &GetPrizePoolLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取活动奖品池，某活动中的奖品信息
func (l *GetPrizePoolLogic) GetPrizePool(in *luckydraw_pb.GetPrizePoolReq) (*luckydraw_pb.GetPrizePoolResp, error) {
	res, err := l.svcCtx.PrizePoolModel.FindOne(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		return nil, nil
	}

	return &luckydraw_pb.GetPrizePoolResp{
		PrizePool: dataconv.MdPrizePoolToPbPrizePool(res),
	}, nil
}
