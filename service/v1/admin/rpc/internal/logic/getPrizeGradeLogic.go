package logic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/service/v1/admin/rpc/admin_pb"
	"mixbench/service/v1/admin/rpc/internal/svc"
	"mixbench/service/v1/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPrizeGradeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPrizeGradeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPrizeGradeLogic {
	return &GetPrizeGradeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPrizeGradeLogic) GetPrizeGrade(in *admin_pb.ItemReq) (*admin_pb.GetPrizeGradeResp, error) {
	res, err := l.svcCtx.PrizeGradeModel.FindOne(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		return nil, nil
	}

	return &admin_pb.GetPrizeGradeResp{
		PrizeGrade: mdPrizeGradeToPbPrizeGrade(res),
	}, nil
}
