package prizesPool

import (
	"context"

	"mixbench/common/respx"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
	"mixbench/service/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PrizesPoolPutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPrizesPoolPutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PrizesPoolPutLogic {
	return &PrizesPoolPutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PrizesPoolPutLogic) PrizesPoolPut(req *types.PrizesPoolPutReq) (any, error) {
	_, err := l.svcCtx.AdminRpc.PutPrizesPool(l.ctx, &admin_pb.PrizesPool{
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
