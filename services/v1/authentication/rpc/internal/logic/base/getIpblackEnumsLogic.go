package baselogic

import (
	"context"

	"mixbench/services/v1/authentication/rpc/authentication_pb"
	"mixbench/services/v1/authentication/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetIpblackEnumsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetIpblackEnumsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetIpblackEnumsLogic {
	return &GetIpblackEnumsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取IP黑名单枚举列表
func (l *GetIpblackEnumsLogic) GetIpblackEnums(in *authentication_pb.GetIpblackEnumsReq) (*authentication_pb.Enums, error) {
	// todo: add your logic here and delete this line

	return &authentication_pb.Enums{}, nil
}
