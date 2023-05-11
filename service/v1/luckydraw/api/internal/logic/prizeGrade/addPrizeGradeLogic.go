package prizeGrade

import (
	"context"

	"mixbench/common/respx"
	"mixbench/service/v1/luckydraw/api/internal/svc"
	"mixbench/service/v1/luckydraw/api/internal/types"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPrizeGradeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddPrizeGradeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPrizeGradeLogic {
	return &AddPrizeGradeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddPrizeGradeLogic) AddPrizeGrade(req *types.AddPrizeGradeReq) (any, error) {
	in := &luckydraw_pb.AddPrizeGradeReq{
		Name:  req.Name,
		IsWin: req.IsWin,
	}
	_, err := l.svcCtx.LuckydrawRpc.AddPrizeGrade(l.ctx, in)
	if err != nil {
		//若rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}

	return respx.DefaultStateResp(nil)
}
