package ipblack

import (
	"context"

	"mixbench/common/respx"
	"mixbench/services/v1/authentication/api/internal/svc"
	"mixbench/services/v1/authentication/api/internal/types"
	"mixbench/services/v1/authentication/rpc/authentication_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelIpblackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelIpblackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelIpblackLogic {
	return &DelIpblackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelIpblackLogic) DelIpblack(req *types.DelIpblackReq) (any, error) {
	_, err := l.svcCtx.AuthenticationRpc.DelIpblack(l.ctx, &authentication_pb.DelIpblackReq{
		Id: req.Id,
	})
	if err != nil {
		//若rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}

	return respx.DefaultStatusResp(nil)
}
