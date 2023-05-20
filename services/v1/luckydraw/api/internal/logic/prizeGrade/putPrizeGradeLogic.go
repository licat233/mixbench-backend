package prizeGrade

import (
	"context"

	"mixbench/common/respx"
	"mixbench/services/v1/luckydraw/api/internal/svc"
	"mixbench/services/v1/luckydraw/api/internal/types"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutPrizeGradeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPutPrizeGradeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutPrizeGradeLogic {
	return &PutPrizeGradeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutPrizeGradeLogic) PutPrizeGrade(req *types.PutPrizeGradeReq) (any, error) {
	in := &luckydraw_pb.PutPrizeGradeReq{
		Id:    req.Id,
		Uuid:  req.Uuid,
		Name:  req.Name,
		IsWin: req.IsWin,
	}
	_, err := l.svcCtx.LuckydrawRpc.PutPrizeGrade(l.ctx, in)
	if err != nil {
		//若rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}

	return respx.DefaultStatusResp(nil)
}
