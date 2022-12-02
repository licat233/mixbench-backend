package prizeGrade

import (
	"context"

	"mixbench/common/respx"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
	"mixbench/service/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PrizeGradeDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPrizeGradeDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PrizeGradeDelLogic {
	return &PrizeGradeDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PrizeGradeDelLogic) PrizeGradeDel(req *types.PrizeGradeDelReq) (any, error) {
	_, err := l.svcCtx.AdminRpc.DelPrizeGrade(l.ctx, &admin_pb.ItemReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return respx.DefaultStateResp(err)
}
