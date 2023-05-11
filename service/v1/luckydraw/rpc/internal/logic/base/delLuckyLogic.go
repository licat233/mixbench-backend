package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/service/v1/luckydraw/rpc/internal/svc"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelLuckyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelLuckyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelLuckyLogic {
	return &DelLuckyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除中奖名单
func (l *DelLuckyLogic) DelLucky(in *luckydraw_pb.DelLuckyReq) (*luckydraw_pb.DelLuckyResp, error) {
	err := l.svcCtx.LuckyModel.Delete(l.ctx, in.Id)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &luckydraw_pb.DelLuckyResp{}, nil
}
