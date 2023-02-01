package prizesPool

import (
	"context"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
	"mixbench/service/v1/admin/rpc/admin_pb"

	"mixbench/common/respx"

	"github.com/zeromicro/go-zero/core/logx"
)

type PrizePoolDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPrizePoolDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PrizePoolDelLogic {
	return &PrizePoolDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PrizePoolDelLogic) PrizePoolDel(req *types.PrizePoolDelReq) (any, error) {
	_, err := l.svcCtx.AdminRpc.DelPrizePool(l.ctx, &admin_pb.ItemReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return respx.DefaultStateResp(err)
}
