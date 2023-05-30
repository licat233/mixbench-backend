package baselogic

import (
	"context"
	"time"

	"mixbench/common/errorx"
	"mixbench/common/uniqueid"
	"mixbench/common/xcrypt"
	"mixbench/services/v1/admin/model"
	"mixbench/services/v1/admin/rpc/admin_pb"
	"mixbench/services/v1/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddAdminerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddAdminerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddAdminerLogic {
	return &AddAdminerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加管理员
func (l *AddAdminerLogic) AddAdminer(in *admin_pb.AddAdminerReq) (*admin_pb.AddAdminerResp, error) {
	if err := xcrypt.PasswordValidator(in.Password); err != nil {
		return nil, errorx.ExtRpcErr(err)
	}
	pw, err := xcrypt.HashPassword(in.Password)
	if err != nil {
		return nil, errorx.ExtRpcErr(err)
	}
	data := &model.Adminer{
		Uuid:         uniqueid.NewUUID(),
		Name:         in.Name,
		Avatar:       in.Avatar,
		Access:       in.Access,
		Passport:     in.Passport,
		Password:     pw,
		Email:        in.Email,
		Status:       in.Status,
		IsSuperAdmin: in.IsSuperAdmin,
		LoginCount:   in.LoginCount,
		LastLogin:    time.Unix(in.LastLogin, 0).Local(),
	}

	_, err = l.svcCtx.AdminerModel.Insert(l.ctx, data)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.AddAdminerResp{}, nil
}
