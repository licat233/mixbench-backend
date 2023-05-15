package prizeGrade

import (
	"context"

	"mixbench/common/respx"
	"mixbench/services/v1/luckydraw/api/internal/svc"
	"mixbench/services/v1/luckydraw/api/internal/types"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPrizeGradeEnumsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPrizeGradeEnumsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPrizeGradeEnumsLogic {
	return &GetPrizeGradeEnumsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPrizeGradeEnumsLogic) GetPrizeGradeEnums(req *types.GetPrizeGradeEnumsReq) (any, error) {
	rpcResp, err := l.svcCtx.LuckydrawRpc.GetPrizeGradeEnums(l.ctx, &luckydraw_pb.GetPrizeGradeEnumsReq{
		ParentId: req.ParentId,
	})
	if err != nil {
		return nil, err
	}
	data := make([]*types.Enum, 0)
	for _, v := range rpcResp.Enums {
		data = append(data, &types.Enum{
			Label: v.Label,
			Value: v.Value,
		})
	}

	return respx.DefaultSingleResp(data, nil)
}
