package prizePool

import (
	"context"

	"mixbench/common/respx"
	"mixbench/service/v1/luckydraw/api/internal/logic/dataconv"
	"mixbench/service/v1/luckydraw/api/internal/svc"
	"mixbench/service/v1/luckydraw/api/internal/types"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPrizePoolLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPrizePoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPrizePoolLogic {
	return &GetPrizePoolLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPrizePoolLogic) GetPrizePool(req *types.GetPrizePoolReq) (any, error) {
	rpcResp, err := l.svcCtx.LuckydrawRpc.GetPrizePool(l.ctx, &luckydraw_pb.GetPrizePoolReq{
		Id: req.Id,
	})
	if err != nil {
		//若rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}
	pbPrizePool := rpcResp.PrizePool
	data := dataconv.PbPrizePoolToApiPrizePool(pbPrizePool)

	return respx.DefaultSingleResp(data, nil)
}
