package ipblack

import (
	"context"

	"mixbench/common/respx"
	"mixbench/services/v1/authentication/api/internal/logic/dataconv"
	"mixbench/services/v1/authentication/api/internal/svc"
	"mixbench/services/v1/authentication/api/internal/types"
	"mixbench/services/v1/authentication/rpc/authentication_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetIpblackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetIpblackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetIpblackLogic {
	return &GetIpblackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetIpblackLogic) GetIpblack(req *types.GetIpblackReq) (any, error) {
	rpcResp, err := l.svcCtx.AuthenticationRpc.GetIpblack(l.ctx, &authentication_pb.GetIpblackReq{
		Id: req.Id,
	})
	if err != nil {
		//若rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}
	pbIpblack := rpcResp.Ipblack
	data := dataconv.PbIpblackToApiIpblack(pbIpblack)

	return respx.DefaultSingleResp(data, nil)
}
