package adminer

import (
	"context"
	"net/http"
	"strings"

	"mixbench/common/errorx"
	"mixbench/common/jwtx"
	"mixbench/common/respx"
	"mixbench/services/v1/admin/api/internal/svc"
	"mixbench/services/v1/admin/api/internal/types"
	"mixbench/services/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewVerifyPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *VerifyPasswordLogic {
	return &VerifyPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *VerifyPasswordLogic) VerifyPassword(req *types.VerifyPasswordReq) (any, error) {
	id, err := jwtx.GetIdByRequest(l.r, "")
	if err != nil {
		return nil, errorx.InternalError(err)
	}
	resp, err := l.svcCtx.AdminRpc.VerifyPassword(l.ctx, &admin_pb.VerifyPasswordReq{
		AdminerId: id,
		Password:  strings.TrimSpace(req.Password),
	})
	if err != nil {
		return nil, err
	}
	return respx.StatusResp("", resp.Ok), nil
}
