package logic

import (
	"context"
	"net/http"

	"mixbench/common/errorx"
	"mixbench/services/v1/admin/api/internal/svc"
	"mixbench/services/v1/admin/rpc/admin_pb"

	"github.com/gorilla/sessions"
	"github.com/zeromicro/go-zero/core/logx"
)

type DevPanelAuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	w      http.ResponseWriter
	r      *http.Request
}

func NewDevPanelAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext, w http.ResponseWriter, r *http.Request) *DevPanelAuthLogic {
	return &DevPanelAuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		w:      w,
		r:      r,
	}
}

const sessionAuthKey = "session-dev"

var store = sessions.NewCookieStore([]byte("secret-key"))

func (l *DevPanelAuthLogic) DevPanelAuth() error {
	//优先调用session校验
	session, _ := store.Get(l.r, sessionAuthKey)

	// 检查用户是否已经登录
	auth, ok := session.Values["authenticated"].(bool)
	if ok && auth {
		return nil
	}

	cookie, err := l.r.Cookie(cookieAuthName)
	if err != nil && err != http.ErrNoCookie {
		// 处理读取 Cookie 失败的情况
		l.Logger.Error(err)
		return errorx.InternalError(err)
	}
	if err == http.ErrNoCookie {
		return errorx.ErrAuth
	}

	cookieValue := cookie.Value
	resp, err := l.svcCtx.AdminRpc.ValidateToken(l.ctx, &admin_pb.ValidateTokenReq{
		JwtToken: cookieValue,
	})
	if err != nil {
		return err
	}
	if !resp.Ok {
		return errorx.ErrAuth
	}

	return nil
}
