package logic

import (
	"context"
	"net/http"

	"mixbench/common/errorx"
	"mixbench/common/utils"
	"mixbench/services/v1/chat/api/internal/svc"
	"mixbench/services/v1/chat/api/internal/wshub"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
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

func (l *VisitorServerLogic) VisitorServer() error {
	if !utils.IsWebsocketRequest(l.r) {
		err := errorx.ExternalError("the current request is not a WebSocket request")
		httpx.ErrorCtx(l.r.Context(), l.w, err)
		return nil
	}

	// 创建一个visitor client
	cli, err := wshub.NewVisitorClient(l.ctx, wshub.HubSvc, l.w, l.r)
	if err != nil {
		return errorx.ExternalError(err)
	}

	if cli == nil {
		//初始化失败
		return nil
	}

	// 开始收发ws信息
	err = cli.Run()
	if err != nil {
		return errorx.InternalError(err)
	}
	return nil
}
