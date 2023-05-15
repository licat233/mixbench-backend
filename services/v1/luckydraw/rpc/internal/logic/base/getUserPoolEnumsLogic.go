package baselogic

import (
	"context"

	"mixbench/services/v1/luckydraw/rpc/internal/svc"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserPoolEnumsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserPoolEnumsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserPoolEnumsLogic {
	return &GetUserPoolEnumsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取活动用户池，可以参与该活动抽奖的用户，及其在该活动中的积分管理枚举列表
func (l *GetUserPoolEnumsLogic) GetUserPoolEnums(in *luckydraw_pb.GetUserPoolEnumsReq) (*luckydraw_pb.Enums, error) {
	// todo: add your logic here and delete this line

	return &luckydraw_pb.Enums{}, nil
}
