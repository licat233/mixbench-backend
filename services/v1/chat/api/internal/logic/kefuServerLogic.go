package logic

import (
	"context"
	"net/http"

	"mixbench/common/ctxdata"
	"mixbench/common/errorx"
	"mixbench/common/utils"
	"mixbench/services/v1/chat/api/internal/logic/dataconv"
	"mixbench/services/v1/chat/api/internal/svc"
	"mixbench/services/v1/chat/api/internal/wshub"
	"mixbench/services/v1/chat/model"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
)

type KefuServerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	w      http.ResponseWriter
	r      *http.Request
	id     int64
	conn   *websocket.Conn
}

func NewKefuServerLogic(ctx context.Context, svcCtx *svc.ServiceContext, w http.ResponseWriter, r *http.Request) *KefuServerLogic {
	return &KefuServerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		w:      w,
		r:      r,
		id:     0,
		conn:   nil,
	}
}

func (l *KefuServerLogic) KefuServer() error {
	if !utils.IsWebsocketRequest(l.r) {
		return errorx.ExternalError("the current request is not a WebSocket request")
	}
	// 获取 id
	id := ctxdata.GetInt64FromCtx(l.ctx, "id")

	kefuM, err := l.svcCtx.KefuModel.FindOne(l.ctx, id)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return errorx.InternalError(err)
	}

	kefuA := dataconv.MdKefuToApiKefu(kefuM)

	// 创建一个kefu client
	cli, err := wshub.NewKefuClient(l.ctx, wshub.HubSvc, l.w, l.r, kefuA)
	if err != nil {
		//可能因为r不是ws请求，导致的报错
		return errorx.ExternalError(err)
	}

	// 开始收发ws信息
	err = cli.Run()
	if err != nil {
		return errorx.InternalError(err)
	}
	return nil
}
