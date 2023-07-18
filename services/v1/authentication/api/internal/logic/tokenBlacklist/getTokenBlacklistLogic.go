package tokenBlacklist

import (
	"context"

	"mixbench/common/respx"
	"mixbench/services/v1/authentication/api/internal/logic/dataconv"
	"mixbench/services/v1/authentication/api/internal/svc"
	"mixbench/services/v1/authentication/api/internal/types"
	"mixbench/services/v1/authentication/rpc/authentication_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTokenBlacklistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTokenBlacklistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTokenBlacklistLogic {
	return &GetTokenBlacklistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTokenBlacklistLogic) GetTokenBlacklist(req *types.GetTokenBlacklistReq) (any, error) {
	rpcResp, err := l.svcCtx.AuthenticationRpc.GetTokenBlacklist(l.ctx, &authentication_pb.GetTokenBlacklistReq{
		Id: req.Id,
	})
	if err != nil {
		//若rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}
	pbTokenBlacklist := rpcResp.TokenBlacklist
	data := dataconv.PbTokenBlacklistToApiTokenBlacklist(pbTokenBlacklist)

	return respx.DefaultSingleResp(data, nil)
}
