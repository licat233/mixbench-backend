package prizesPool

import (
	"context"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
	"mixbench/service/v1/admin/rpc/admin_pb"

	"mixbench/common/respx"

	"github.com/zeromicro/go-zero/core/logx"
)

type PrizesPoolInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPrizesPoolInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PrizesPoolInfoLogic {
	return &PrizesPoolInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PrizesPoolInfoLogic) PrizesPoolInfo(req *types.PrizesPoolGetReq) (any, error) {
	resp, err := l.svcCtx.AdminRpc.GetPrizesPool(l.ctx, &admin_pb.ItemReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return respx.DefaultSingleResp(resp.PrizesPool, err)
}
