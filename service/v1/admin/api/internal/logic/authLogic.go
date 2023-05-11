package logic

import (
	"context"
	"net/http"

	"mixbench/service/v1/admin/api/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthLogic struct {
	logx.Logger
	ctx     context.Context
	svcCtx  *svc.ServiceContext
	Request *http.Request
}

func NewAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *AuthLogic {
	return &AuthLogic{
		Logger:  logx.WithContext(ctx),
		ctx:     ctx,
		svcCtx:  svcCtx,
		Request: r,
	}
}

func (l *AuthLogic) Auth() error {
	return nil
}
