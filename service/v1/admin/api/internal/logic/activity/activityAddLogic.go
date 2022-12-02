package activity

import (
	"context"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
	"mixbench/service/v1/admin/rpc/admin_pb"

	"mixbench/common/errorx"
	"mixbench/common/respx"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActivityAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewActivityAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActivityAddLogic {
	return &ActivityAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ActivityAddLogic) ActivityAdd(req *types.ActivityAddReq) (any, error) {
	if req.Name == "" {
		return nil, errorx.ErrInvalidParam
	}

	resp, err := l.svcCtx.AdminRpc.AddActivity(l.ctx, &admin_pb.Activity{
		Name:      req.Name,
		ProjectId: req.ProjectId,
		Status:    0,
		Pattern:   req.Pattern,
		MaxCount:  req.MaxCount,
		MaxPrizes: req.MaxPrizes,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	})
	if err != nil {
		return nil, err
	}
	return respx.DefaultSingleResp(resp.Activity, err)
}
