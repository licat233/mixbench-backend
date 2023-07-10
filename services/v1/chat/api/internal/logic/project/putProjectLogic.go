package project

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/common/respx"
	"mixbench/services/v1/chat/api/internal/svc"
	"mixbench/services/v1/chat/api/internal/types"
	"mixbench/services/v1/chat/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPutProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutProjectLogic {
	return &PutProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutProjectLogic) PutProject(req *types.PutProjectReq) (any, error) {
	in := &model.Project{
		Id:     req.Id,
		Uuid:   req.Uuid,
		Name:   req.Name,
		Status: req.Status,
	}
	if err := l.svcCtx.ProjectModel.Update(l.ctx, in); err != nil {
		l.Logger.Error("failed to update project, error: ", err)
		return nil, errorx.InternalError(err)
	}

	return respx.DefaultStatusResp(nil)
}
