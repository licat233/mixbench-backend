package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/service/v1/luckydraw/rpc/internal/svc"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelPrizeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelPrizeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelPrizeLogic {
	return &DelPrizeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除奖品名单
func (l *DelPrizeLogic) DelPrize(in *luckydraw_pb.DelPrizeReq) (*luckydraw_pb.DelPrizeResp, error) {
	err := l.svcCtx.PrizeModel.Delete(l.ctx, in.Id)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &luckydraw_pb.DelPrizeResp{}, nil
}
