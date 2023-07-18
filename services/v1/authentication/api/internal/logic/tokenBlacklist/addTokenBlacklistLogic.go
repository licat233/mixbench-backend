package tokenBlacklist

import (
	"context"

	"mixbench/common/respx"
	"mixbench/services/v1/authentication/api/internal/svc"
	"mixbench/services/v1/authentication/api/internal/types"
	"mixbench/services/v1/authentication/rpc/authentication_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTokenBlacklistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddTokenBlacklistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTokenBlacklistLogic {
	return &AddTokenBlacklistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddTokenBlacklistLogic) AddTokenBlacklist(req *types.AddTokenBlacklistReq) (any, error) {
	in := &authentication_pb.AddTokenBlacklistReq{
		Token:    req.Token,
		Platform: req.Platform,
		Ip:       req.Ip,
	}
	if _, err := l.svcCtx.AuthenticationRpc.AddTokenBlacklist(l.ctx, in); err != nil {
		//若rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}

	return respx.DefaultStatusResp(nil)
}
