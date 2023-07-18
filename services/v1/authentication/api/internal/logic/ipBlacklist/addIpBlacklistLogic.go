package ipBlacklist

import (
	"context"

	"mixbench/common/respx"
	"mixbench/services/v1/authentication/api/internal/svc"
	"mixbench/services/v1/authentication/api/internal/types"
	"mixbench/services/v1/authentication/rpc/authentication_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddIpBlacklistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddIpBlacklistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddIpBlacklistLogic {
	return &AddIpBlacklistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddIpBlacklistLogic) AddIpBlacklist(req *types.AddIpBlacklistReq) (any, error) {
	in := &authentication_pb.AddIpBlacklistReq{
		Ip: req.Ip,
	}
	if _, err := l.svcCtx.AuthenticationRpc.AddIpBlacklist(l.ctx, in); err != nil {
		//若rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}

	return respx.DefaultStatusResp(nil)
}
