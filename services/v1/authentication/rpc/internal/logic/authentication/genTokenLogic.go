package authenticationlogic

import (
	"context"

	"mixbench/services/v1/authentication/rpc/authentication_pb"
	"mixbench/services/v1/authentication/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenTokenLogic {
	return &GenTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 生成jwt-token
func (l *GenTokenLogic) GenToken(in *authentication_pb.GenTokenReq) (*authentication_pb.GenTokenResp, error) {
	// todo: add your logic here and delete this line

	return &authentication_pb.GenTokenResp{}, nil
}
