package logic

import (
	"context"
	"errors"
	"net/http"

	"mixbench/services/v1/admin/api/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	req    *http.Request
	resp   http.ResponseWriter
	err    error
}

func NewAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext, w http.ResponseWriter, r *http.Request) *AuthLogic {
	return &AuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		req:    r,
		resp:   w,
		err:    errors.New("authentication failed"),
	}
}

func (l *AuthLogic) Auth() error {
	// jwtToken := l.req.Header.Get(jwtx.ReqHeaderKey)
	// jwtToken = strings.TrimSpace(jwtToken)
	// if jwtToken == "" {
	// 	return l.err
	// }
	return nil
}
