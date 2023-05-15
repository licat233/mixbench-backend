package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/common/uniqueid"
	"mixbench/services/v1/luckydraw/model"
	"mixbench/services/v1/luckydraw/rpc/internal/svc"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserPoolLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserPoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserPoolLogic {
	return &AddUserPoolLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加活动用户池，可以参与该活动抽奖的用户，及其在该活动中的积分管理
func (l *AddUserPoolLogic) AddUserPool(in *luckydraw_pb.AddUserPoolReq) (*luckydraw_pb.AddUserPoolResp, error) {
	data := &model.UserPool{
		Uuid:       uniqueid.NewUUID(),
		ActivityId: in.ActivityId,
		UserId:     in.UserId,
		Status:     in.Status,
		Score:      in.Score,
	}
	_, err := l.svcCtx.UserPoolModel.Insert(l.ctx, data)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &luckydraw_pb.AddUserPoolResp{}, nil
}
