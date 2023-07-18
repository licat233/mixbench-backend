package ipBlacklist

import (
	"context"

	"mixbench/services/v1/authentication/api/internal/svc"
	"mixbench/services/v1/authentication/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetIpBlacklistEnumsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetIpBlacklistEnumsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetIpBlacklistEnumsLogic {
	return &GetIpBlacklistEnumsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetIpBlacklistEnumsLogic) GetIpBlacklistEnums(req *types.GetIpBlacklistEnumsReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
