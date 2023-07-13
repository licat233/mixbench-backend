package web

import (
	"context"

	"mixbench/common/respx"
	"mixbench/services/v1/admin/api/internal/logic/dataconv"
	"mixbench/services/v1/admin/api/internal/svc"
	"mixbench/services/v1/admin/api/internal/types"
	"mixbench/services/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetWebLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetWebLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWebLogic {
	return &GetWebLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetWebLogic) GetWeb(req *types.GetWebReq) (any, error) {
	rpcResp, err := l.svcCtx.AdminRpc.GetWeb(l.ctx, &admin_pb.GetWebReq{
		Id: req.Id,
	})
	if err != nil {
		//若rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}
	pbWeb := rpcResp.Web
	data := dataconv.PbWebToApiWeb(pbWeb)

	return respx.DefaultSingleResp(data, nil)
}
