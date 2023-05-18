package adminlogic

import (
	"context"

	"mixbench/services/v1/admin/rpc/admin_pb"
	"mixbench/services/v1/admin/rpc/internal/svc"

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

// 预留的调试端口
func (l *TestLogic) Test(in *admin_pb.TestReq) (*admin_pb.TestResp, error) {
	return &admin_pb.TestResp{
		Data: in.Data,
	}, nil
}
