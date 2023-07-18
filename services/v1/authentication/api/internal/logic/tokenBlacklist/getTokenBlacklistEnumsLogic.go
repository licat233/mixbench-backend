package tokenBlacklist

import (
	"context"

	"mixbench/services/v1/authentication/api/internal/svc"
	"mixbench/services/v1/authentication/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTokenBlacklistEnumsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTokenBlacklistEnumsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTokenBlacklistEnumsLogic {
	return &GetTokenBlacklistEnumsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTokenBlacklistEnumsLogic) GetTokenBlacklistEnums(req *types.GetTokenBlacklistEnumsReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
