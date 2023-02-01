/*
 * @Author: licat
 * @Date: 2023-01-31 17:18:24
 * @LastEditors: licat
 * @LastEditTime: 2023-01-31 17:33:12
 * @Description: licat233@gmail.com
 */
package logic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/service/v1/luckydraw/rpc/internal/logic/tools"
	"mixbench/service/v1/luckydraw/rpc/internal/svc"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type BaseGetPrizeGradeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBaseGetPrizeGradeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BaseGetPrizeGradeListLogic {
	return &BaseGetPrizeGradeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BaseGetPrizeGradeListLogic) BaseGetPrizeGradeList(in *luckydraw_pb.GetPrizeGradeListReq) (*luckydraw_pb.GetPrizeGradeListResp, error) {
	in.ListReq = tools.NewListReq(in.ListReq)
	list, total, err := l.svcCtx.PrizeGradeModel.FindList(l.ctx, in.ListReq.PageSize, in.ListReq.Page, in.ListReq.Keyword, tools.PbPrizeGradeToMdPrizeGrade(in.PrizeGrade))
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &luckydraw_pb.GetPrizeGradeListResp{
		PrizeGrades: tools.MdPrizeGradesToPbPrizeGrades(list),
		Total:       total,
	}, nil
}
