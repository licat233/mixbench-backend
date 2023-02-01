/*
 * @Author: licat
 * @Date: 2023-01-31 17:18:24
 * @LastEditors: licat
 * @LastEditTime: 2023-01-31 17:38:35
 * @Description: licat233@gmail.com
 */
package logic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/service/v1/luckydraw/model"
	"mixbench/service/v1/luckydraw/rpc/internal/svc"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type BasePutPrizeGradeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBasePutPrizeGradeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BasePutPrizeGradeLogic {
	return &BasePutPrizeGradeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BasePutPrizeGradeLogic) BasePutPrizeGrade(in *luckydraw_pb.PutPrizeGradeReq) (*luckydraw_pb.PutPrizeGradeResp, error) {
	data := &model.PrizeGrade{
		Id:    in.Id,
		Name:  in.Name,
		IsWin: in.IsWin,
	}

	if err := l.svcCtx.PrizeGradeModel.Update(l.ctx, data); err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	return &luckydraw_pb.PutPrizeGradeResp{}, nil
}
