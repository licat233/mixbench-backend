package logic

import (
	"context"
	"mixbench/service/v1/admin/rpc/admin_pb"
	"mixbench/service/v1/admin/rpc/internal/svc"
	"mixbench/service/v1/model"

	"mixbench/common/errorx"
	"mixbench/common/utils"

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

func (l *PutActivityLogic) PutActivity(in *admin_pb.Activity) (*admin_pb.StatusResp, error) {
	//是否为空
	if in == nil {
		return nil, errorx.SecRpcErr(errorx.ErrEmptyParam)
	}

	//构建model.activity
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
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return newStatusResponse(), nil
}
