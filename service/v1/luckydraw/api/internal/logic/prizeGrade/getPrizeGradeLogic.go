package prizeGrade

import (
	"context"

	"mixbench/common/respx"
	"mixbench/service/v1/luckydraw/api/internal/logic/dataconv"
	"mixbench/service/v1/luckydraw/api/internal/svc"
	"mixbench/service/v1/luckydraw/api/internal/types"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPrizeGradeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPrizeGradeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPrizeGradeLogic {
	return &GetPrizeGradeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPrizeGradeLogic) GetPrizeGrade(req *types.GetPrizeGradeReq) (any, error) {
	rpcResp, err := l.svcCtx.LuckydrawRpc.GetPrizeGrade(l.ctx, &luckydraw_pb.GetPrizeGradeReq{
		Id: req.Id,
	})
	if err != nil {
		//若rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}
	pbPrizeGrade := rpcResp.PrizeGrade
	data := dataconv.PbPrizeGradeToApiPrizeGrade(pbPrizeGrade)

	return respx.DefaultSingleResp(data, nil)
}
