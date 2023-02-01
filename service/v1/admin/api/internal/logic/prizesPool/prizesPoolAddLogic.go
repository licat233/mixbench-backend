package prizesPool

import (
	"context"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
	"mixbench/service/v1/admin/rpc/admin_pb"

	"mixbench/common/respx"

	"github.com/zeromicro/go-zero/core/logx"
)

type PrizePoolAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPrizePoolAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PrizePoolAddLogic {
	return &PrizePoolAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PrizePoolAddLogic) PrizePoolAdd(req *types.PrizePoolAddReq) (any, error) {
	resp, err := l.svcCtx.AdminRpc.AddPrizePool(l.ctx, &admin_pb.PrizePool{
		ActivityId: req.ActivityId,
		PrizeId:    req.PrizeId,
		GradeId:    req.GradeId,
		LuckValue:  req.LuckValue,
		Status:     req.Status,
		Total:      req.Total,
		Remainder:  req.Remainder,
	})
	if err != nil {
		return nil, err
	}
	return respx.DefaultSingleResp(resp.PrizePool, err)
}
