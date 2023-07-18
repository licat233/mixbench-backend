package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/services/v1/luckydraw/rpc/internal/logic/dataconv"
	"mixbench/services/v1/luckydraw/rpc/internal/svc"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutPrizeGradeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPutPrizeGradeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutPrizeGradeLogic {
	return &PutPrizeGradeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新奖品级别
func (l *PutPrizeGradeLogic) PutPrizeGrade(in *luckydraw_pb.PutPrizeGradeReq) (*luckydraw_pb.PutPrizeGradeResp, error) {
	data := dataconv.PbPrizeGradeToMdPrizeGrade(in.PrizeGrade)
	if err := l.svcCtx.PrizeGradeModel.Update(l.ctx, data); err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &luckydraw_pb.PutPrizeGradeResp{}, nil
}
