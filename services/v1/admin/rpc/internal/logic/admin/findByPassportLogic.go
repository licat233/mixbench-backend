package adminlogic

import (
	"context"
	"errors"

	"mixbench/common/errorx"
	"mixbench/services/v1/admin/model"
	"mixbench/services/v1/admin/rpc/admin_pb"
	"mixbench/services/v1/admin/rpc/internal/logic/dataconv"
	"mixbench/services/v1/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindByPassportLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindByPassportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindByPassportLogic {
	return &FindByPassportLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 通过账号查找管理员
func (l *FindByPassportLogic) FindByPassport(in *admin_pb.FindByPassportReq) (*admin_pb.FindByPassportResp, error) {
	if len(in.Passport) == 0 {
		return nil, errorx.ExtRpcErr(errors.New("passport is empty"))
	}
	adminer, err := l.svcCtx.AdminerModel.FindByPassport(l.ctx, in.Passport)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.FindByPassportResp{
		Adminer: dataconv.MdAdminerToPbAdminer(adminer),
	}, nil
}
