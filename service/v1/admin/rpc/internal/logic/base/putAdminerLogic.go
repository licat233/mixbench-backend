package baselogic

import (
	"context"
	"time"

	"mixbench/common/errorx"
	"mixbench/service/v1/admin/model"
	"mixbench/service/v1/admin/rpc/admin_pb"
	"mixbench/service/v1/admin/rpc/internal/svc"

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

// 更新管理员
func (l *PutAdminerLogic) PutAdminer(in *admin_pb.PutAdminerReq) (*admin_pb.PutAdminerResp, error) {
	data := &model.Adminer{
		Id:           in.Id,
		Uuid:         in.Uuid,
		Name:         in.Name,
		Avatar:       in.Avatar,
		Access:       in.Access,
		Passport:     in.Passport,
		Password:     in.Password,
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
