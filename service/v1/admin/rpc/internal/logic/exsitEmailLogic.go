package logic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/service/v1/admin/rpc/admin_pb"
	"mixbench/service/v1/admin/rpc/internal/svc"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
)

type ExsitEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewExsitEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExsitEmailLogic {
	return &ExsitEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 邮箱是否存在
func (l *ExsitEmailLogic) ExsitEmail(in *admin_pb.ExsitEmailReq) (*admin_pb.ExsitEmailResp, error) {
	//校验邮箱
	if !govalidator.IsEmail(in.Email) {
		return nil, errorx.ExtRpcErr(errorx.ErrInvalidEmail)
	}
	//查找是否有绑定该邮箱的用户
	exist, err := l.svcCtx.AdminerModel.ExsitByEmail(l.ctx, in.Email)
	if err != nil {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	return &admin_pb.ExsitEmailResp{
		Exsit: exist,
	}, nil
}
