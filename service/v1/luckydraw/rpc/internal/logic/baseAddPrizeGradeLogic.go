/*
 * @Author: licat
 * @Date: 2023-01-31 17:18:24
 * @LastEditors: licat
 * @LastEditTime: 2023-01-31 17:24:06
 * @Description: licat233@gmail.com
 */
package logic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/service/v1/luckydraw/model"
	"mixbench/service/v1/luckydraw/rpc/internal/logic/tools"
	"mixbench/service/v1/luckydraw/rpc/internal/svc"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type BaseAddPrizeGradeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBaseAddPrizeGradeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BaseAddPrizeGradeLogic {
	return &BaseAddPrizeGradeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BaseAddPrizeGradeLogic) BaseAddPrizeGrade(in *luckydraw_pb.AddPrizeGradeReq) (*luckydraw_pb.AddPrizeGradeResp, error) {
	data := &model.PrizeGrade{
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

	return &luckydraw_pb.AddPrizeGradeResp{
		PrizeGrade: tools.MdPrizeGradeToPbPrizeGrade(prizeGrade),
	}, nil
}
