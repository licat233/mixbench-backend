package prizesPool

import (
	"context"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
	"mixbench/service/v1/admin/rpc/admin_pb"

	"mixbench/common/respx"

	"github.com/zeromicro/go-zero/core/logx"
)

type PrizePoolInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPrizePoolInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PrizePoolInfoLogic {
	return &PrizePoolInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PrizePoolInfoLogic) PrizePoolInfo(req *types.PrizePoolGetReq) (any, error) {
	resp, err := l.svcCtx.AdminRpc.GetPrizePool(l.ctx, &admin_pb.ItemReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return respx.DefaultSingleResp(resp.PrizePool, err)
}
