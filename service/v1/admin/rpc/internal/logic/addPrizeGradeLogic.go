package logic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/service/v1/admin/rpc/admin_pb"
	"mixbench/service/v1/admin/rpc/internal/svc"
	"mixbench/service/v1/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPrizeGradeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddPrizeGradeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPrizeGradeLogic {
	return &AddPrizeGradeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddPrizeGradeLogic) AddPrizeGrade(in *admin_pb.PrizeGrade) (*admin_pb.AddPrizeGradeResp, error) {
	if in == nil {
		return nil, errorx.SecRpcErr(errorx.ErrEmptyParam)
	}
	data := &model.PrizeGrade{
		Id:    in.Id,
		Name:  in.Name,
		IsWin: in.IsWin,
	}
	result, err := l.svcCtx.PrizeGradeModel.Insert(l.ctx, data)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	prizeGrade, err := l.svcCtx.PrizeGradeModel.FindOne(l.ctx, id)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.AddPrizeGradeResp{
		PrizeGrade: mdPrizeGradeToPbPrizeGrade(prizeGrade),
	}, nil
}
