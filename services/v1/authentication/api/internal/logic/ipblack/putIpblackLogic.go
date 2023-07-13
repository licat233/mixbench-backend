package ipblack

import (
	"context"

	"mixbench/common/respx"
	"mixbench/services/v1/authentication/api/internal/svc"
	"mixbench/services/v1/authentication/api/internal/types"
	"mixbench/services/v1/authentication/rpc/authentication_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutIpblackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPutIpblackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutIpblackLogic {
	return &PutIpblackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutIpblackLogic) PutIpblack(req *types.PutIpblackReq) (any, error) {
	in := &authentication_pb.PutIpblackReq{
		Id: req.Id,
		Ip: req.Ip,
	}
	if _, err := l.svcCtx.AuthenticationRpc.PutIpblack(l.ctx, in); err != nil {
		//若rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}

	return respx.DefaultStatusResp(nil)
}
