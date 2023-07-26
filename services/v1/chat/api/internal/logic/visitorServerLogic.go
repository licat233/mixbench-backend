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

func (l *VisitorServerLogic) VisitorServer(req *types.ConnWebsocketReq) (_ error) {
	//鉴权
	token := req.Token
	if token = strings.TrimSpace(token); token == "" {
		l.w.WriteHeader(http.StatusUnauthorized)
		return
	}
	claims, ok := jwtx.ValidateToken(token, l.svcCtx.Config.VisitorAuthSecret)
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

	//查找游客是否存在
	if id < 1 {
		//游客不存在
		l.w.WriteHeader(http.StatusUnauthorized)
		return
	}
	visitor, err := l.svcCtx.VisitorModel.FindById(l.ctx, id)
	if err == model.ErrNotFound {
		//游客不存在
		l.w.WriteHeader(http.StatusUnauthorized)
		return
	} else if err != nil {
		l.Logger.Error(err)
		l.w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//如果该游客已被拉黑 status == 2
	if visitor.Status != 1 && visitor.Status != 0 {
		l.w.WriteHeader(http.StatusUnauthorized)
		return
	}

	//查找项目是否存在
	if visitor.ProjectId < 1 {
		l.w.WriteHeader(http.StatusUnauthorized)
		return
	}
	projectResp, err := l.svcCtx.AdminRpc.GetProject(l.ctx, &admin_pb.GetProjectReq{
		Id: visitor.ProjectId,
	})
	if err != nil {
		l.w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if projectResp.Project == nil {
		//不存在
		l.w.WriteHeader(http.StatusUnauthorized)
		return
	}
	//检查状态
	if projectResp.Project.Status != 1 {
		l.w.WriteHeader(http.StatusUnauthorized)
		return
	}

	//至此，鉴权结束

	cli, err := ws.NewVisitorClient(l.ctx, l.svcCtx, l.w, l.r, &ws.VisitorInfo{Visitor: visitor})
	if err != nil {
		l.Logger.Error(err)
		l.w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 开始收发ws信息
	cli.Run()
	return
}
