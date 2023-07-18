package lucky

import (
	"context"

	"mixbench/common/respx"
	"mixbench/services/v1/luckydraw/api/internal/svc"
	"mixbench/services/v1/luckydraw/api/internal/types"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutLuckyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPutLuckyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutLuckyLogic {
	return &PutLuckyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutLuckyLogic) PutLucky(req *types.PutLuckyReq) (any, error) {
	in := &luckydraw_pb.PutLuckyReq{
		Lucky: &luckydraw_pb.Lucky{
			Id:          req.Id,
			Uuid:        req.Uuid,
			UserId:      req.UserId,
			PrizePoolId: req.PrizePoolId,
			Status:      req.Status,
			LuckDate:    req.LuckDate,
			Ip:          req.Ip,
		},
	}
	if _, err := l.svcCtx.LuckydrawRpc.PutLucky(l.ctx, in); err != nil {
		//若rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}

	return respx.DefaultStatusResp(nil)
}
