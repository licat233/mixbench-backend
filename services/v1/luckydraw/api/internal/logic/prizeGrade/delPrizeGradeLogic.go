package prizeGrade

import (
	"context"

	"mixbench/common/respx"
	"mixbench/services/v1/luckydraw/api/internal/svc"
	"mixbench/services/v1/luckydraw/api/internal/types"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelPrizeGradeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelPrizeGradeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelPrizeGradeLogic {
	return &DelPrizeGradeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelPrizeGradeLogic) DelPrizeGrade(req *types.DelPrizeGradeReq) (any, error) {
	_, err := l.svcCtx.LuckydrawRpc.DelPrizeGrade(l.ctx, &luckydraw_pb.DelPrizeGradeReq{
		Id: req.Id,
	})
	if err != nil {
		//若rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}

	return respx.DefaultStateResp(nil)
}
