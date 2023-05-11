package prize

import (
	"context"

	"mixbench/common/respx"
	"mixbench/service/v1/luckydraw/api/internal/svc"
	"mixbench/service/v1/luckydraw/api/internal/types"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelPrizeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelPrizeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelPrizeLogic {
	return &DelPrizeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelPrizeLogic) DelPrize(req *types.DelPrizeReq) (any, error) {
	_, err := l.svcCtx.LuckydrawRpc.DelPrize(l.ctx, &luckydraw_pb.DelPrizeReq{
		Id: req.Id,
	})
	if err != nil {
		//若rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}

	return respx.DefaultStateResp(nil)
}
