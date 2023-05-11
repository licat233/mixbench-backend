package baselogic

import (
	"context"

	"mixbench/service/v1/luckydraw/rpc/internal/svc"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPrizePoolEnumsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPrizePoolEnumsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPrizePoolEnumsLogic {
	return &GetPrizePoolEnumsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取活动奖品池，某活动中的奖品信息枚举列表
func (l *GetPrizePoolEnumsLogic) GetPrizePoolEnums(in *luckydraw_pb.GetPrizePoolEnumsReq) (*luckydraw_pb.Enums, error) {
	// todo: add your logic here and delete this line

	return &luckydraw_pb.Enums{}, nil
}
