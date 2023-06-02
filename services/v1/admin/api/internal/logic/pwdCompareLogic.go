package logic

import (
	"context"

	"mixbench/common/xcrypt"
	"mixbench/services/v1/admin/api/internal/svc"
	"mixbench/services/v1/admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PwdCompareLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPwdCompareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PwdCompareLogic {
	return &PwdCompareLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PwdCompareLogic) PwdCompare(req *types.PwdCompareReq) (any, error) {
	ok := xcrypt.CheckPassword(req.Password, req.Hash)
	return ok, nil
}
