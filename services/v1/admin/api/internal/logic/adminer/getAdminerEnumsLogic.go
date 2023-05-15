package adminer

import (
	"context"

	"mixbench/common/respx"
	"mixbench/services/v1/admin/api/internal/svc"
	"mixbench/services/v1/admin/api/internal/types"
	"mixbench/services/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAdminerEnumsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAdminerEnumsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdminerEnumsLogic {
	return &GetAdminerEnumsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAdminerEnumsLogic) GetAdminerEnums(req *types.GetAdminerEnumsReq) (any, error) {
	rpcResp, err := l.svcCtx.AdminRpc.GetAdminerEnums(l.ctx, &admin_pb.GetAdminerEnumsReq{
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
