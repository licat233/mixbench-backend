package prize

import (
	"context"

	"mixbench/common/respx"
	"mixbench/services/v1/luckydraw/api/internal/svc"
	"mixbench/services/v1/luckydraw/api/internal/types"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetActivityPrizesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetActivityPrizesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetActivityPrizesLogic {
	return &GetActivityPrizesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetActivityPrizesLogic) GetActivityPrizes(req *types.GetActivityPrizesReq) (any, error) {
	in := &luckydraw_pb.GetActivityPrizesReq{
		ActivityUuid: "",
		ActivityId:   req.ActivityId,
	}
	res, err := l.svcCtx.LuckydrawRpc.GetActivityPrizes(l.ctx, in)
	if err != nil {
		return nil, err
	}
	return respx.DefaultSingleResp(res.ActivityPrizes, err)
}
