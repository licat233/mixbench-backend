package prizeGrade

import (
	"context"

	"mixbench/common/respx"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
	"mixbench/service/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PrizeGradeInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPrizeGradeInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PrizeGradeInfoLogic {
	return &PrizeGradeInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PrizeGradeInfoLogic) PrizeGradeInfo(req *types.PrizeGradeGetReq) (any, error) {
	resp, err := l.svcCtx.AdminRpc.GetPrizeGrade(l.ctx, &admin_pb.ItemReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return respx.DefaultSingleResp(resp.PrizeGrade, err)
}
