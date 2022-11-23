package rbac

import (
	"context"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutPoliciesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPutPoliciesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutPoliciesLogic {
	return &PutPoliciesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutPoliciesLogic) PutPolicies(req *types.PoliciesReq) (any, error) {
	//TODO: 待开发
	// l.svcCtx.CasbinEnforcer.UpdatePolicies()
	return nil, nil
}
