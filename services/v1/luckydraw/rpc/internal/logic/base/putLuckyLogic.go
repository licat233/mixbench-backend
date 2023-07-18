package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/services/v1/luckydraw/rpc/internal/logic/dataconv"
	"mixbench/services/v1/luckydraw/rpc/internal/svc"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutLuckyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPutLuckyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutLuckyLogic {
	return &PutLuckyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新中奖名单
func (l *PutLuckyLogic) PutLucky(in *luckydraw_pb.PutLuckyReq) (*luckydraw_pb.PutLuckyResp, error) {
	data := dataconv.PbLuckyToMdLucky(in.Lucky)
	if err := l.svcCtx.LuckyModel.Update(l.ctx, data); err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &luckydraw_pb.PutLuckyResp{}, nil
}
