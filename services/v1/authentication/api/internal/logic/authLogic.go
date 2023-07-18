package logic

import (
	"context"
	"errors"
	"net/http"

	"mixbench/services/v1/authentication/api/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	w      http.ResponseWriter
	r      *http.Request
}

func NewAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext, w http.ResponseWriter, r *http.Request) *AuthLogic {
	return &AuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		w:      w,
		r:      r,
	}
}

func (l *AuthLogic) Auth() error {
	if err := l.svcCtx.AuthMiddlewareLogic.AuthHandler(l.w, l.r); err != nil {
		return errors.New("authentication failed")
	}
	return nil
}
