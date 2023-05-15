package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/services/v1/luckydraw/model"
	"mixbench/services/v1/luckydraw/rpc/internal/svc"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutUserPoolLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPutUserPoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutUserPoolLogic {
	return &PutUserPoolLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新活动用户池，可以参与该活动抽奖的用户，及其在该活动中的积分管理
func (l *PutUserPoolLogic) PutUserPool(in *luckydraw_pb.PutUserPoolReq) (*luckydraw_pb.PutUserPoolResp, error) {
	data := &model.UserPool{
		Id:         in.Id,
		Uuid:       in.Uuid,
		ActivityId: in.ActivityId,
		UserId:     in.UserId,
		Status:     in.Status,
		Score:      in.Score,
	}

	if err := l.svcCtx.UserPoolModel.Update(l.ctx, data); err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &luckydraw_pb.PutUserPoolResp{}, nil
}
