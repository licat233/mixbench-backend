package activity

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/common/respx"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
	"mixbench/service/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActivityPutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewActivityPutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActivityPutLogic {
	return &ActivityPutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ActivityPutLogic) ActivityPut(req *types.ActivityPutReq) (any, error) {
	if req.Name == "" {
		return nil, errorx.ErrInvalidParam
	}
	_, err := l.svcCtx.AdminRpc.PutActivity(l.ctx, &admin_pb.Activity{
		Id:        req.Id,
		Uuid:      req.Uuid,
		Name:      req.Name,
		ProjectId: req.ProjectId,
		Status:    req.Status,
		Pattern:   req.Pattern,
		MaxCount:  req.MaxCount,
		MaxPrizes: req.MaxPrizes,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	})
	if err != nil {
		return nil, err
	}
	return respx.DefaultStateResp(err)
}
