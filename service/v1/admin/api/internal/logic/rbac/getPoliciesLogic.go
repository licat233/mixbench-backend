package rbac

import (
	"context"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"mixbench/common/respx"
)

type GetPoliciesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPoliciesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPoliciesLogic {
	return &GetPoliciesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPoliciesLogic) GetPolicies(req *types.PoliciesReq) (any, error) {
	return respx.DefaultSingleResp(l.svcCtx.CasbinEnforcer.GetPolicy(), nil)
}
