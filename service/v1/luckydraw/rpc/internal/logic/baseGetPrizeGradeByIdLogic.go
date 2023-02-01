/*
 * @Author: licat
 * @Date: 2023-01-31 17:18:24
 * @LastEditors: licat
 * @LastEditTime: 2023-01-31 17:32:42
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

type BaseGetPrizeGradeByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBaseGetPrizeGradeByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BaseGetPrizeGradeByIdLogic {
	return &BaseGetPrizeGradeByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BaseGetPrizeGradeByIdLogic) BaseGetPrizeGradeById(in *luckydraw_pb.GetPrizeGradeByIdReq) (*luckydraw_pb.GetPrizeGradeByIdResp, error) {
	data, err := l.svcCtx.PrizeGradeModel.FindOne(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	return &luckydraw_pb.GetPrizeGradeByIdResp{
		PrizeGrade: tools.MdPrizeGradeToPbPrizeGrade(data),
	}, nil
}
