package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/service/v1/luckydraw/rpc/internal/svc"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelActivityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelActivityLogic {
	return &DelActivityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除活动
func (l *DelActivityLogic) DelActivity(in *luckydraw_pb.DelActivityReq) (*luckydraw_pb.DelActivityResp, error) {
	err := l.svcCtx.ActivityModel.Delete(l.ctx, in.Id)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &luckydraw_pb.DelActivityResp{}, nil
}
