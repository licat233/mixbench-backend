package lucky

import (
	"context"

	"mixbench/common/respx"
	"mixbench/service/v1/luckydraw/api/internal/svc"
	"mixbench/service/v1/luckydraw/api/internal/types"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLuckyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLuckyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLuckyLogic {
	return &AddLuckyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLuckyLogic) AddLucky(req *types.AddLuckyReq) (any, error) {
	in := &luckydraw_pb.AddLuckyReq{
		UserId:      req.UserId,
		PrizePoolId: req.PrizePoolId,
		Status:      req.Status,
		LuckDate:    req.LuckDate,
		Ip:          req.Ip,
	}
	_, err := l.svcCtx.LuckydrawRpc.AddLucky(l.ctx, in)
	if err != nil {
		//若rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}

	return respx.DefaultStateResp(nil)
}
