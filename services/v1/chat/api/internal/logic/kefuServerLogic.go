package logic

import (
	"context"
	"net/http"
	"strings"

	"mixbench/common/jwtx"
	"mixbench/services/v1/admin/rpc/admin_pb"
	"mixbench/services/v1/chat/api/internal/svc"
	"mixbench/services/v1/chat/api/internal/types"
	"mixbench/services/v1/chat/api/internal/ws"

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

func (l *KefuServerLogic) KefuServer(req *types.ConnWebsocketReq) (_ error) {
	//鉴权
	token := req.Token
	if token = strings.TrimSpace(token); token == "" {
		l.w.WriteHeader(http.StatusUnauthorized)
		return nil
	}
	claims, ok := jwtx.ValidateToken(string(token), l.svcCtx.Config.VisitorAuthSecret)
	if !ok {
		//鉴权失败
		l.w.WriteHeader(http.StatusUnauthorized)
		return
	}

	id, ok := jwtx.GetId(claims)
	if !ok {
		//鉴权失败
		l.w.WriteHeader(http.StatusUnauthorized)
		return
	}

	//查找客服是否存在
	adminerResp, err := l.svcCtx.AdminRpc.GetAdminer(l.ctx, &admin_pb.GetAdminerReq{Id: id})
	if err != nil {
		// l.Logger.Error(err)
		l.w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if adminerResp.Adminer == nil {
		//不存在
		l.w.WriteHeader(http.StatusUnauthorized)
		return
	}

	kefu := adminerResp.Adminer

	//查找项目是否存在
	projectResp, err := l.svcCtx.AdminRpc.GetProject(l.ctx, &admin_pb.GetProjectReq{Id: kefu.ProjectId})
	if err != nil {
		// l.Logger.Error(err)
		l.w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if projectResp.Project == nil {
		//不存在
		l.w.WriteHeader(http.StatusUnauthorized)
		return
	}

	cli, err := ws.NewKefuClient(l.ctx, l.svcCtx, l.w, l.r, &ws.KefuInfo{Adminer: kefu})
	if err != nil {
		l.Logger.Error(err)
		l.w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 开始收发ws信息
	cli.Run()
	return
}
