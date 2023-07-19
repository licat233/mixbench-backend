package logic

import (
	"context"

	"mixbench/services/v1/authentication/rpc/authentication_pb"
	"mixbench/services/v1/authentication/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthLogic {
	return &AuthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// jwt鉴权
func (l *AuthLogic) Auth(in *authentication_pb.AuthReq) (*authentication_pb.AuthResp, error) {
	// todo: add your logic here and delete this line

	return &authentication_pb.AuthResp{}, nil
}
