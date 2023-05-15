package prize

import (
	"context"

	"mixbench/common/respx"
	"mixbench/services/v1/luckydraw/api/internal/logic/dataconv"
	"mixbench/services/v1/luckydraw/api/internal/svc"
	"mixbench/services/v1/luckydraw/api/internal/types"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPrizeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPrizeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPrizeLogic {
	return &GetPrizeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPrizeLogic) GetPrize(req *types.GetPrizeReq) (any, error) {
	rpcResp, err := l.svcCtx.LuckydrawRpc.GetPrize(l.ctx, &luckydraw_pb.GetPrizeReq{
		Id: req.Id,
	})
	if err != nil {
		//若rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}
	pbPrize := rpcResp.Prize
	data := dataconv.PbPrizeToApiPrize(pbPrize)

	return respx.DefaultSingleResp(data, nil)
}
