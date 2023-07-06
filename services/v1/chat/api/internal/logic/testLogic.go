package logic

import (
	"context"
	"net/http"

	"mixbench/common/ctxdata"
	"mixbench/common/errorx"
	"mixbench/common/jwtx"
	"mixbench/common/respx"
	"mixbench/services/v1/chat/api/internal/svc"
	"mixbench/services/v1/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewTestLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *TestLogic {
	return &TestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *TestLogic) Test(req *types.TestReq) (any, error) {
	token := jwtx.GetToken(l.r)
	claims, err := jwtx.GetClaims(l.r, l.svcCtx.Config.Auth.AccessSecret)
	if err != nil {
		return nil, errorx.InternalError(err)
	}
	data := map[string]any{
		"param":         req.Data,
		"token_value":   token,
		"token_details": claims,
		"cxt_id":        ctxdata.GetUidFromCtx(l.ctx),
	}

	return respx.NewData(data), nil
}
