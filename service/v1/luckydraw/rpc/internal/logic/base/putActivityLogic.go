package baselogic

import (
	"context"
	"time"

	"mixbench/common/errorx"
	"mixbench/service/v1/luckydraw/model"
	"mixbench/service/v1/luckydraw/rpc/internal/svc"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

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
	data := &model.Activity{
		Id:        in.Id,
		Uuid:      in.Uuid,
		Name:      in.Name,
		ProjectId: in.ProjectId,
		Status:    in.Status,
		Pattern:   in.Pattern,
		MaxCount:  in.MaxCount,
		MaxPrizes: in.MaxPrizes,
		StartTime: time.Unix(in.StartTime, 0).Local(),
		EndTime:   time.Unix(in.EndTime, 0).Local(),
	}

	if err := l.svcCtx.ActivityModel.Update(l.ctx, data); err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &luckydraw_pb.PutActivityResp{}, nil
}
