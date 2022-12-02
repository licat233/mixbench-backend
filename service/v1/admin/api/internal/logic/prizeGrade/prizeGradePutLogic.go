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

type PrizeGradePutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPrizeGradePutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PrizeGradePutLogic {
	return &PrizeGradePutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PrizeGradePutLogic) PrizeGradePut(req *types.PrizeGradePutReq) (any, error) {
	if req.Name == "" {
		return nil, errorx.ErrInvalidParam
	}
	_, err := l.svcCtx.AdminRpc.PutPrizeGrade(l.ctx, &admin_pb.PrizeGrade{
		Id:    req.Id,
		Name:  req.Name,
		IsWin: req.IsWin,
	})
	if err != nil {
		return nil, err
	}
	return respx.DefaultStateResp(err)
}
