package prizesPool

import (
	"context"

	"mixbench/common/respx"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
	"mixbench/service/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PrizePoolPutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPrizePoolPutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PrizePoolPutLogic {
	return &PrizePoolPutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PrizePoolPutLogic) PrizePoolPut(req *types.PrizePoolPutReq) (any, error) {
	_, err := l.svcCtx.AdminRpc.PutPrizePool(l.ctx, &admin_pb.PrizePool{
		Id:         req.Id,
		Uuid:       req.Uuid,
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
	return respx.DefaultStateResp(err)
}
