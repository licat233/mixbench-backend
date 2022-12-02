package adminer

import (
	"context"
	"errors"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
	"mixbench/service/v1/admin/rpc/admin_pb"

	"mixbench/common/errorx"
	"mixbench/common/respx"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminerAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminerAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminerAddLogic {
	return &AdminerAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminerAddLogic) AdminerAdd(req *types.AdminerAddReq) (any, error) {
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
	resp, err := l.svcCtx.AdminRpc.AddAdminer(l.ctx, &admin_pb.Adminer{
		Name:         req.Name,
		Avatar:       req.Avatar,
		Passport:     req.Passport,
		Password:     req.Password,
		Email:        req.Email,
		Status:       req.Status,
		IsSuperAdmin: req.IsSuperAdmin,
	})
	if err != nil {
		return nil, err
	}
	return respx.DefaultSingleResp(resp.Adminer, err)
}
