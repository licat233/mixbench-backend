package ipBlacklist

import (
	"context"

	"mixbench/common/respx"
	"mixbench/services/v1/authentication/api/internal/logic/dataconv"
	"mixbench/services/v1/authentication/api/internal/svc"
	"mixbench/services/v1/authentication/api/internal/types"
	"mixbench/services/v1/authentication/rpc/authentication_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetIpBlacklistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetIpBlacklistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetIpBlacklistLogic {
	return &GetIpBlacklistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetIpBlacklistLogic) GetIpBlacklist(req *types.GetIpBlacklistReq) (any, error) {
	rpcResp, err := l.svcCtx.AuthenticationRpc.GetIpBlacklist(l.ctx, &authentication_pb.GetIpBlacklistReq{
		Id: req.Id,
	})
	if err != nil {
		//若rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}
	pbIpBlacklist := rpcResp.IpBlacklist
	data := dataconv.PbIpBlacklistToApiIpBlacklist(pbIpBlacklist)

	return respx.DefaultSingleResp(data, nil)
}
