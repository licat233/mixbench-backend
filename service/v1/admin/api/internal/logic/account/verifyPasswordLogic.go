package account

import (
	"context"
	"strings"

	"mixbench/common/ctxdata"
	"mixbench/common/respx"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
	"mixbench/service/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVerifyPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyPasswordLogic {
	return &VerifyPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerifyPasswordLogic) VerifyPassword(req *types.VerifyPasswordReq) (any, error) {
	uid := ctxdata.GetUidFromCtx(l.ctx)
	resp, err := l.svcCtx.AdminRpc.VerifyPassword(l.ctx, &admin_pb.VerifyPasswordReq{
		AdminerId: uid,
		Password:  strings.TrimSpace(req.Password),
	})
	if err != nil {
		return nil, err
	}
	return respx.StateResp("", resp.Ok), nil
}
