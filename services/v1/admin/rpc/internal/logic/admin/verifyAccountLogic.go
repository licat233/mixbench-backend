package adminlogic

import (
	"context"

	"mixbench/services/v1/admin/rpc/admin_pb"
	"mixbench/services/v1/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyAccountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVerifyAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyAccountLogic {
	return &VerifyAccountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 生产token
func (l *VerifyAccountLogic) VerifyAccount(in *admin_pb.VerifyAccountReq) (*admin_pb.VerifyAccountResp, error) {
	// todo: add your logic here and delete this line

	return &admin_pb.VerifyAccountResp{}, nil
}
