package user

import (
	"context"

	"mixbench/common/respx"
	"mixbench/services/v1/luckydraw/api/internal/svc"
	"mixbench/services/v1/luckydraw/api/internal/types"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserEnumsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserEnumsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserEnumsLogic {
	return &GetUserEnumsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserEnumsLogic) GetUserEnums(req *types.GetUserEnumsReq) (any, error) {
	rpcResp, err := l.svcCtx.LuckydrawRpc.GetUserEnums(l.ctx, &luckydraw_pb.GetUserEnumsReq{
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
