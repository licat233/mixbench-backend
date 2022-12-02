package prizeGrade

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/common/respx"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
	"mixbench/service/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PrizeGradeAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPrizeGradeAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PrizeGradeAddLogic {
	return &PrizeGradeAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PrizeGradeAddLogic) PrizeGradeAdd(req *types.PrizeGradeAddReq) (any, error) {
	if req.Name == "" {
		return nil, errorx.ErrInvalidParam
	}
	resp, err := l.svcCtx.AdminRpc.AddPrizeGrade(l.ctx, &admin_pb.PrizeGrade{
		Name:  req.Name,
		IsWin: req.IsWin,
	})
	if err != nil {
		return nil, err
	}
	return respx.DefaultSingleResp(resp.PrizeGrade, err)
}
