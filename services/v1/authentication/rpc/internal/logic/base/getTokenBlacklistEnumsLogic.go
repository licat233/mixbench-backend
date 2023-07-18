package baselogic

import (
	"context"

	"mixbench/services/v1/authentication/rpc/authentication_pb"
	"mixbench/services/v1/authentication/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTokenBlacklistEnumsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTokenBlacklistEnumsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTokenBlacklistEnumsLogic {
	return &GetTokenBlacklistEnumsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取token黑名单枚举列表
func (l *GetTokenBlacklistEnumsLogic) GetTokenBlacklistEnums(in *authentication_pb.GetTokenBlacklistEnumsReq) (*authentication_pb.Enums, error) {
	// todo: add your logic here and delete this line

	return &authentication_pb.Enums{}, nil
}
