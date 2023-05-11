package lucky

import (
	"context"

	"mixbench/common/respx"
	"mixbench/service/v1/luckydraw/api/internal/svc"
	"mixbench/service/v1/luckydraw/api/internal/types"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelLuckyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelLuckyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelLuckyLogic {
	return &DelLuckyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelLuckyLogic) DelLucky(req *types.DelLuckyReq) (any, error) {
	_, err := l.svcCtx.LuckydrawRpc.DelLucky(l.ctx, &luckydraw_pb.DelLuckyReq{
		Id: req.Id,
	})
	if err != nil {
		//若rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}

	return respx.DefaultStateResp(nil)
}
