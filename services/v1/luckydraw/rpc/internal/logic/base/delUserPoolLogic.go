package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/services/v1/luckydraw/rpc/internal/svc"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelUserPoolLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelUserPoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelUserPoolLogic {
	return &DelUserPoolLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除活动用户池，可以参与该活动抽奖的用户，及其在该活动中的积分管理
func (l *DelUserPoolLogic) DelUserPool(in *luckydraw_pb.DelUserPoolReq) (*luckydraw_pb.DelUserPoolResp, error) {
	err := l.svcCtx.UserPoolModel.SoftDelete(l.ctx, in.Id)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &luckydraw_pb.DelUserPoolResp{}, nil
}
