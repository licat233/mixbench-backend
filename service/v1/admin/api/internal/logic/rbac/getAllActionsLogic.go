package rbac

import (
	"context"
	"mixbench/service/v1/admin/api/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"mixbench/common/respx"
)

type GetAllActionsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllActionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllActionsLogic {
	return &GetAllActionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllActionsLogic) GetAllActions() (any, error) {
	data := l.svcCtx.CasbinEnforcer.GetAllActions()
	return respx.DefaultSingleResp(data, nil)
}
