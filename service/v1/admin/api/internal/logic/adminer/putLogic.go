package adminer

import (
	"context"
	"errors"

	"mixbench/common/errorx"
	"mixbench/common/respx"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
	"mixbench/service/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutLogic {
	return &PutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutLogic) Put(req *types.AdminerPutReq) (any, error) {
	ok, err := SuperAdminerIdentify(l.ctx, l.svcCtx)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errorx.AccessError(nil, errors.New("no request permission"))
	}
	if req.Name == "" {
		return nil, errorx.ErrInvalidParam
	}
	_, err = l.svcCtx.AdminRpc.PutAdminer(l.ctx, &admin_pb.Adminer{
		Id:           req.Id,
		Uuid:         req.Uuid,
		Name:         req.Name,
		Avatar:       req.Avatar,
		Passport:     req.Passport,
		Password:     req.Password,
		Email:        req.Email,
		Status:       req.Status,
		IsSuperAdmin: req.IsSuperAdmin,
		LoginCount:   req.LoginCount,
		LastLogin:    req.LastLogin,
	})
	if err != nil {
		return nil, err
	}
	return respx.DefaultStateResp(err)
}
