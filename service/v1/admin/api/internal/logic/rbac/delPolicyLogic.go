package rbac

import (
	"context"
	"errors"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"

	"mixbench/common/errorx"
	"mixbench/common/respx"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelPolicyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelPolicyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelPolicyLogic {
	return &DelPolicyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelPolicyLogic) DelPolicy(req *types.PolicyReq) (any, error) {
	ok, err := l.svcCtx.CasbinEnforcer.RemovePolicy(req.Sub, req.Obj, req.Act)
	if logState && err != nil {
		logx.Error(err)
		return nil, errorx.InternalError(err)
	}
	return respx.DefaultBoolResp(ok, errors.New("rule not exists"))
}
