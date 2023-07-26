package logic

import (
	"context"
	"net/http"
	"strings"

	"mixbench/services/v1/chat/api/internal/svc"
	"mixbench/services/v1/chat/api/internal/types"
	"mixbench/services/v1/chat/api/internal/ws"

	"github.com/zeromicro/go-zero/core/logx"
)

type VisitorServerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	w      http.ResponseWriter
	r      *http.Request
}

func NewVisitorServerLogic(ctx context.Context, svcCtx *svc.ServiceContext, w http.ResponseWriter, r *http.Request) *VisitorServerLogic {
	return &VisitorServerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		w:      w,
		r:      r,
	}
}

func (l *VisitorServerLogic) VisitorServer(req *types.ConnWebsocketReq) (_ error) {
	//鉴权
	token := req.Token
	if token = strings.TrimSpace(token); token == "" {
		l.w.WriteHeader(http.StatusUnauthorized)
		return
	}
	login := NewVisitorLoginLogic(l.ctx, l.svcCtx, l.r)
	token, visitor, err := login.Validator(token)
	if err != nil {
		l.w.WriteHeader(http.StatusUnauthorized)
		return
	}

	//至此，鉴权结束

	cli, err := ws.NewVisitorClient(l.ctx, l.svcCtx, l.w, l.r, &ws.VisitorInfo{Visitor: visitor}, token)
	if err != nil {
		l.Logger.Error(err)
		l.w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 开始收发ws信息
	cli.Run()
	return
}
