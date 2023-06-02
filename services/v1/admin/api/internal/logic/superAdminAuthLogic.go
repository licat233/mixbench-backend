package logic

import (
	"context"
	"errors"
	"net/http"

	"mixbench/services/v1/admin/api/internal/middleware"
	"mixbench/services/v1/admin/api/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SuperAdminAuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	req    *http.Request
	err    error
}

func NewSuperAdminAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *SuperAdminAuthLogic {
	return &SuperAdminAuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		req:    r,
		err:    errors.New("authentication failed"),
	}
}

func (l *SuperAdminAuthLogic) SuperAdminAuth() error {
	adminer, err := middleware.AuthLogic(l.ctx, l.req, l.svcCtx.Config, l.svcCtx.AdminerModel, l.svcCtx.JwtBlacklistModel)
	if err != nil {
		return l.err
	}
	if adminer.IsSuperAdmin == 0 {
		return l.err
	}
	return nil
}
