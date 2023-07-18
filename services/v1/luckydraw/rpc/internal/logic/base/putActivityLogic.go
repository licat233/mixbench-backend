package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/services/v1/luckydraw/rpc/internal/logic/dataconv"
	"mixbench/services/v1/luckydraw/rpc/internal/svc"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutActivityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPutActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutActivityLogic {
	return &PutActivityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新活动
func (l *PutActivityLogic) PutActivity(in *luckydraw_pb.PutActivityReq) (*luckydraw_pb.PutActivityResp, error) {
	data := dataconv.PbActivityToMdActivity(in.Activity)
	if err := l.svcCtx.ActivityModel.Update(l.ctx, data); err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &luckydraw_pb.PutActivityResp{}, nil
}
