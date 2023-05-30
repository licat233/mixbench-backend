package adminlogic

import (
	"context"
	"errors"

	"mixbench/common/errorx"
	"mixbench/services/v1/admin/rpc/admin_pb"
	"mixbench/services/v1/admin/rpc/internal/svc"
	"mixbench/services/v1/luckydraw/model"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
)

type ForgotPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewForgotPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ForgotPasswordLogic {
	return &ForgotPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 忘记密码，通过邮箱重置密码
func (l *ForgotPasswordLogic) ForgotPassword(in *admin_pb.ForgotPasswordReq) (*admin_pb.StatusResp, error) {
	//邮箱格式验证
	if !govalidator.IsEmail(in.Email) {
		return nil, errorx.ExtRpcErr(errors.New("invalid email"))
	}
	//查找是否存在绑定该邮箱的adminer
	_, err := l.svcCtx.AdminerModel.FindByEmail(l.ctx, in.Email)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		return nil, errorx.ExtRpcErr(errors.New("the adminer does not exist"))
	}
	return &admin_pb.StatusResp{
		Status: true,
	}, nil
}
