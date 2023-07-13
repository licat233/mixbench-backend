package authenticationlogic

import (
	"context"

	"mixbench/services/v1/authentication/rpc/authentication_pb"
	"mixbench/services/v1/authentication/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestLogic {
	return &TestLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Preset methods can be deleted. At least one method needs to be defined in a service
func (l *TestLogic) Test(in *authentication_pb.NilReq) (*authentication_pb.NilResp, error) {
	// todo: add your logic here and delete this line

	return &authentication_pb.NilResp{}, nil
}
