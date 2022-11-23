package rbac

import (
	"context"
	"mixbench/service/v1/admin/api/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"mixbench/common/respx"
)

type GetAllPolicyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllPolicyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllPolicyLogic {
	return &GetAllPolicyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllPolicyLogic) GetAllPolicy() (any, error) {
	return respx.DefaultSingleResp(l.svcCtx.CasbinEnforcer.GetPolicy(), nil)
}
