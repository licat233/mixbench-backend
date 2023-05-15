package baselogic

import (
	"context"

	"mixbench/services/v1/luckydraw/rpc/internal/svc"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLuckyEnumsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLuckyEnumsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLuckyEnumsLogic {
	return &GetLuckyEnumsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取中奖名单枚举列表
func (l *GetLuckyEnumsLogic) GetLuckyEnums(in *luckydraw_pb.GetLuckyEnumsReq) (*luckydraw_pb.Enums, error) {
	// todo: add your logic here and delete this line

	return &luckydraw_pb.Enums{}, nil
}
