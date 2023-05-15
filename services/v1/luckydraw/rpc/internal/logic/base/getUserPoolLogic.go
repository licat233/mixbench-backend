package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/services/v1/luckydraw/model"
	"mixbench/services/v1/luckydraw/rpc/internal/logic/dataconv"
	"mixbench/services/v1/luckydraw/rpc/internal/svc"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserPoolLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserPoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserPoolLogic {
	return &GetUserPoolLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取活动用户池，可以参与该活动抽奖的用户，及其在该活动中的积分管理
func (l *GetUserPoolLogic) GetUserPool(in *luckydraw_pb.GetUserPoolReq) (*luckydraw_pb.GetUserPoolResp, error) {
	res, err := l.svcCtx.UserPoolModel.FindOne(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		return nil, nil
	}

	return &luckydraw_pb.GetUserPoolResp{
		UserPool: dataconv.MdUserPoolToPbUserPool(res),
	}, nil
}
