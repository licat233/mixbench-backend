package baselogic

import (
	"context"

	"mixbench/services/v1/admin/rpc/admin_pb"
	"mixbench/services/v1/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetWebEnumsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetWebEnumsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWebEnumsLogic {
	return &GetWebEnumsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取网站枚举列表
func (l *GetWebEnumsLogic) GetWebEnums(in *admin_pb.GetWebEnumsReq) (*admin_pb.Enums, error) {
	// todo: add your logic here and delete this line

	return &admin_pb.Enums{}, nil
}
