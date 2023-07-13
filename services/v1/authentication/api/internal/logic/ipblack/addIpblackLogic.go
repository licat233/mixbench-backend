package ipblack

import (
	"context"

	"mixbench/common/respx"
	"mixbench/services/v1/authentication/api/internal/svc"
	"mixbench/services/v1/authentication/api/internal/types"
	"mixbench/services/v1/authentication/rpc/authentication_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddIpblackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddIpblackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddIpblackLogic {
	return &AddIpblackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddIpblackLogic) AddIpblack(req *types.AddIpblackReq) (any, error) {
	in := &authentication_pb.AddIpblackReq{
		Ip: req.Ip,
	}
	if _, err := l.svcCtx.AuthenticationRpc.AddIpblack(l.ctx, in); err != nil {
		//若rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}

	return respx.DefaultStatusResp(nil)
}
