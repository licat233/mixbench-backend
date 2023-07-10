package project

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/common/respx"
	"mixbench/services/v1/chat/api/internal/svc"
	"mixbench/services/v1/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelProjectLogic {
	return &DelProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelProjectLogic) DelProject(req *types.DelProjectReq) (any, error) {
	if err := l.svcCtx.ProjectModel.SoftDelete(l.ctx, req.Id); err != nil {
		l.Logger.Error("failed to soft delete project, error: ", err)
		return nil, errorx.InternalError(err)
	}

	return respx.DefaultStatusResp(nil)
}
