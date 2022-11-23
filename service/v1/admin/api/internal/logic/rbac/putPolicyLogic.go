package rbac

import (
	"context"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutPolicyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPutPolicyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutPolicyLogic {
	return &PutPolicyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutPolicyLogic) PutPolicy(req *types.PolicyReq) (any, error) {
	// todo: 需要自己开发

	return nil, nil
}
