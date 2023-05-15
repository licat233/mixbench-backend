package adminer

import (
	"context"

	"mixbench/common/ctxdata"
	"mixbench/common/respx"
	"mixbench/services/v1/admin/api/internal/svc"
	"mixbench/services/v1/admin/api/internal/types"
	"mixbench/services/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutCurrentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPutCurrentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutCurrentLogic {
	return &PutCurrentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutCurrentLogic) PutCurrent(req *types.PutCurrentReq) (any, error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	resp, err := l.svcCtx.AdminRpc.GetAdminer(l.ctx, &admin_pb.GetAdminerReq{
		Id: userId,
	})
	if err != nil {
		return nil, err
	}

	adminer := resp.Adminer
	adminer.Name = req.Name
	adminer.Avatar = req.Avatar
	adminer.Passport = req.Passport
	adminer.Email = req.Email
	_, err = l.svcCtx.AdminRpc.PutAdminer(l.ctx, &admin_pb.PutAdminerReq{
		Id:           adminer.Id,
		Uuid:         adminer.Uuid,
		Name:         adminer.Name,
		Avatar:       adminer.Avatar,
		Access:       adminer.Access,
		Passport:     adminer.Passport,
		Password:     adminer.Password,
		Email:        adminer.Email,
		Status:       adminer.Status,
		IsSuperAdmin: adminer.IsSuperAdmin,
		LoginCount:   adminer.LoginCount,
		LastLogin:    adminer.LastLogin,
	})
	if err != nil {
		return nil, err
	}
	return respx.DefaultStateResp(err)
}
