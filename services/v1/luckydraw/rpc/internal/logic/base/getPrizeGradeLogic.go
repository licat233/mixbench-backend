package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/services/v1/luckydraw/model"
	"mixbench/services/v1/luckydraw/rpc/internal/logic/dataconv"
	"mixbench/services/v1/luckydraw/rpc/internal/svc"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

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

// 获取奖品级别
func (l *GetPrizeGradeLogic) GetPrizeGrade(in *luckydraw_pb.GetPrizeGradeReq) (*luckydraw_pb.GetPrizeGradeResp, error) {
	res, err := l.svcCtx.PrizeGradeModel.FindById(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &luckydraw_pb.GetPrizeGradeResp{
		PrizeGrade: dataconv.MdPrizeGradeToPbPrizeGrade(res),
	}, nil
}
