package tokenBlacklist

import (
	"context"

	"mixbench/common/respx"
	"mixbench/services/v1/authentication/api/internal/svc"
	"mixbench/services/v1/authentication/api/internal/types"
	"mixbench/services/v1/authentication/rpc/authentication_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutTokenBlacklistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPutTokenBlacklistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutTokenBlacklistLogic {
	return &PutTokenBlacklistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutTokenBlacklistLogic) PutTokenBlacklist(req *types.PutTokenBlacklistReq) (any, error) {
	in := &authentication_pb.PutTokenBlacklistReq{
		TokenBlacklist: &authentication_pb.TokenBlacklist{
			Id:       req.Id,
			Token:    req.Token,
			Platform: req.Platform,
			Ip:       req.Ip,
		},
	}
	if _, err := l.svcCtx.AuthenticationRpc.PutTokenBlacklist(l.ctx, in); err != nil {
		//若rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}

	return respx.DefaultStatusResp(nil)
}
