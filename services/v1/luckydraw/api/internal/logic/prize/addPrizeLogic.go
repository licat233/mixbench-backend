package prize

import (
	"context"

	"mixbench/common/respx"
	"mixbench/services/v1/luckydraw/api/internal/svc"
	"mixbench/services/v1/luckydraw/api/internal/types"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPrizeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddPrizeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPrizeLogic {
	return &AddPrizeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddPrizeLogic) AddPrize(req *types.AddPrizeReq) (any, error) {
	in := &luckydraw_pb.AddPrizeReq{
		Name:   req.Name,
		Image:  req.Image,
		Price:  req.Price,
		Status: req.Status,
	}
	_, err := l.svcCtx.LuckydrawRpc.AddPrize(l.ctx, in)
	if err != nil {
		//若rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}

	return respx.DefaultStatusResp(nil)
}
