package logic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/common/utils"
	"mixbench/service/v1/luckydraw/model"
	"mixbench/service/v1/luckydraw/rpc/internal/svc"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type BasePutActivityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBasePutActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BasePutActivityLogic {
	return &BasePutActivityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BasePutActivityLogic) BasePutActivity(in *luckydraw_pb.PutActivityReq) (*luckydraw_pb.PutActivityResp, error) {
	activity := &model.Activity{
		Id:        in.Id,
		Uuid:      in.Uuid,
		Name:      in.Name,
		ProjectId: in.ProjectId,
		Status:    in.Status,
		Pattern:   in.Pattern,
		MaxCount:  in.MaxCount,
		MaxPrizes: in.MaxPrizes,
		StartTime: utils.Int64ToTime(utils.TimestampHandle(in.StartTime)),
		EndTime:   utils.Int64ToTime(utils.TimestampHandle(in.EndTime)),
	}

	//更新数据
	if err := l.svcCtx.ActivityModel.Update(l.ctx, activity); err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &luckydraw_pb.PutActivityResp{}, nil
}
