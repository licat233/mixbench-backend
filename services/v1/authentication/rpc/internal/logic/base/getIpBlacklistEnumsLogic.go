package baselogic

import (
	"context"

	"mixbench/services/v1/authentication/rpc/authentication_pb"
	"mixbench/services/v1/authentication/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetIpBlacklistEnumsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetIpBlacklistEnumsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetIpBlacklistEnumsLogic {
	return &GetIpBlacklistEnumsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取IP黑名单枚举列表
func (l *GetIpBlacklistEnumsLogic) GetIpBlacklistEnums(in *authentication_pb.GetIpBlacklistEnumsReq) (*authentication_pb.Enums, error) {
	// todo: add your logic here and delete this line

	return &authentication_pb.Enums{}, nil
}
