package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/service/v1/luckydraw/rpc/internal/svc"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelPrizeGradeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelPrizeGradeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelPrizeGradeLogic {
	return &DelPrizeGradeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除奖品级别
func (l *DelPrizeGradeLogic) DelPrizeGrade(in *luckydraw_pb.DelPrizeGradeReq) (*luckydraw_pb.DelPrizeGradeResp, error) {
	err := l.svcCtx.PrizeGradeModel.Delete(l.ctx, in.Id)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &luckydraw_pb.DelPrizeGradeResp{}, nil
}
