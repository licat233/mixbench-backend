package logic

import (
	"context"
	"errors"

	"mixbench/common/errorx"
	"mixbench/service/v1/admin/rpc/admin_pb"
	"mixbench/service/v1/admin/rpc/internal/svc"
	"mixbench/service/v1/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResetLuckValueLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewResetLuckValueLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetLuckValueLogic {
	return &ResetLuckValueLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ResetLuckValueLogic) ResetLuckValue(in *admin_pb.ResetLuckValueReq) (*admin_pb.StatusResp, error) {
	//先验证活动信息是否正确
	activity, err := l.svcCtx.ActivityModel.FindOne(l.ctx, in.ActivityId)
	if err != nil && err != model.ErrNotFound {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		return nil, errorx.ExtRpcErr(errors.New("not found activity"))
	}
	if activity.Uuid != in.ActivityUuid {
		return nil, errorx.ExtRpcErr(errors.New("params invalid"))
	}
	prizesPools, err := l.svcCtx.PrizesPoolModel.FindsByActivityId(l.ctx, in.ActivityId)
	if err != nil {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	for _, prizesPool := range prizesPools {
		prizesPool.LuckValue = 100
		if err = l.svcCtx.PrizesPoolModel.Update(l.ctx, prizesPool); err != nil {
			logx.Error(err)
			return nil, errorx.IntRpcErr(err)
		}
	}
	return &admin_pb.StatusResp{}, nil
}
