package logic

import (
	"context"
	"mixbench/service/v1/admin/rpc/admin_pb"
	"mixbench/service/v1/admin/rpc/internal/svc"

	"mixbench/common/errorx"
	"mixbench/common/jwtx"

	"github.com/zeromicro/go-zero/core/logx"
)

type ParseJwtTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewParseJwtTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ParseJwtTokenLogic {
	return &ParseJwtTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 解析 jwtToken, 返回原生error
func (l *ParseJwtTokenLogic) ParseJwtToken(in *admin_pb.ParseJwtTokenReq) (*admin_pb.ParseJwtTokenResp, error) {
	claims, err := jwtx.ParseToken(in.JwtToken, in.Secret)
	if err != nil {
		return nil, errorx.SecRpcErr(err)
	}
	data := jwtx.ClaimsToMap(*claims)
	return &admin_pb.ParseJwtTokenResp{
		Claims: data,
	}, nil
}
