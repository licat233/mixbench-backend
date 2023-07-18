package tokenBlacklist

import (
	"context"

	"mixbench/common/respx"
	"mixbench/services/v1/authentication/api/internal/svc"
	"mixbench/services/v1/authentication/api/internal/types"
	"mixbench/services/v1/authentication/rpc/authentication_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelTokenBlacklistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelTokenBlacklistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelTokenBlacklistLogic {
	return &DelTokenBlacklistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelTokenBlacklistLogic) DelTokenBlacklist(req *types.DelTokenBlacklistReq) (any, error) {
	_, err := l.svcCtx.AuthenticationRpc.DelTokenBlacklist(l.ctx, &authentication_pb.DelTokenBlacklistReq{
		Id: req.Id,
	})
	if err != nil {
		//若rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}

	return respx.DefaultStatusResp(nil)
}
