package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/services/v1/luckydraw/model"
	"mixbench/services/v1/luckydraw/rpc/internal/logic/dataconv"
	"mixbench/services/v1/luckydraw/rpc/internal/svc"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPrizeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPrizeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPrizeLogic {
	return &GetPrizeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取奖品名单
func (l *GetPrizeLogic) GetPrize(in *luckydraw_pb.GetPrizeReq) (*luckydraw_pb.GetPrizeResp, error) {
	res, err := l.svcCtx.PrizeModel.FindById(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &luckydraw_pb.GetPrizeResp{
		Prize: dataconv.MdPrizeToPbPrize(res),
	}, nil
}
