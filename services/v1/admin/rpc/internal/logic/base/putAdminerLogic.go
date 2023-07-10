package baselogic

import (
	"context"
	"errors"
	"time"

	"mixbench/common/errorx"
	"mixbench/services/v1/admin/model"
	"mixbench/services/v1/admin/rpc/admin_pb"
	"mixbench/services/v1/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutAdminerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPutAdminerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutAdminerLogic {
	return &PutAdminerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新管理员，
func (l *PutAdminerLogic) PutAdminer(in *admin_pb.PutAdminerReq) (*admin_pb.PutAdminerResp, error) {
	// 不更新密码，密码需要通过单独端口进行更新
	adminer, err := l.svcCtx.AdminerModel.FindById(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		return nil, errorx.ExtRpcErr(errors.New("the adminer does not exist"))
	}
	data := &model.Adminer{
		Id:           in.Id,
		Uuid:         in.Uuid,
		Name:         in.Name,
		Avatar:       in.Avatar,
		Access:       in.Access,
		Passport:     in.Passport,
		Password:     adminer.Password,
		Email:        in.Email,
		Status:       in.Status,
		IsSuperAdmin: in.IsSuperAdmin,
		LoginCount:   in.LoginCount,
		LastLogin:    time.Unix(in.LastLogin, 0).Local(),
	}

	if err := l.svcCtx.AdminerModel.Update(l.ctx, data); err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.PutAdminerResp{}, nil
}
