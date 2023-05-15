package project

import (
	"context"

	"mixbench/common/respx"
	"mixbench/services/v1/luckydraw/api/internal/svc"
	"mixbench/services/v1/luckydraw/api/internal/types"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProjectEnumsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProjectEnumsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProjectEnumsLogic {
	return &GetProjectEnumsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProjectEnumsLogic) GetProjectEnums(req *types.GetProjectEnumsReq) (any, error) {
	rpcResp, err := l.svcCtx.LuckydrawRpc.GetProjectEnums(l.ctx, &luckydraw_pb.GetProjectEnumsReq{
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
