package logic

import (
	"context"
	"net/http"

	"mixbench/common/ctxdata"
	"mixbench/common/errorx"
	"mixbench/common/utils"
	"mixbench/services/v1/chat/api/internal/client"
	"mixbench/services/v1/chat/api/internal/logic/dataconv"
	"mixbench/services/v1/chat/api/internal/svc"
	"mixbench/services/v1/chat/model"

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

func (l *VisitorServerLogic) VisitorServer() error {
	if !utils.IsWebsocketRequest(l.r) {
		return errorx.ExternalError("the current request is not a WebSocket request")
	}

	// 获取 id
	id := ctxdata.GetInt64FromCtx(l.ctx, "id")

	kefuM, err := l.svcCtx.VisitorModel.FindOne(l.ctx, id)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return errorx.InternalError(err)
	}

	kefuA := dataconv.MdVisitorToApiVisitor(kefuM)

	// 创建一个visitor client
	cli, err := client.NewVisitorClient(l.ctx, l.svcCtx.Hub, l.w, l.r, kefuA)
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
