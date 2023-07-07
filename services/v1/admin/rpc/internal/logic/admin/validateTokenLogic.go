package adminlogic

import (
	"context"
	"strings"

	"mixbench/common/errorx"
	"mixbench/common/jwtx"
	"mixbench/services/v1/admin/rpc/admin_pb"
	"mixbench/services/v1/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ValidateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewValidateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ValidateTokenLogic {
	return &ValidateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 验证token是否有效
func (l *ValidateTokenLogic) ValidateToken(in *admin_pb.ValidateTokenReq) (*admin_pb.ValidateTokenResp, error) {
	if in.JwtToken = strings.TrimSpace(in.JwtToken); in.JwtToken == "" {
		return &admin_pb.ValidateTokenResp{Ok: false}, nil
	}
	ok, err := jwtx.ValidateToken(in.JwtToken, l.svcCtx.Config.JwtAuth.AccessSecret)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	return &admin_pb.ValidateTokenResp{Ok: ok}, nil
}
