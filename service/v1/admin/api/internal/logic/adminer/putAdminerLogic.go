package adminer

import (
	"context"

	"mixbench/common/respx"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
	"mixbench/service/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutAdminerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPutAdminerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutAdminerLogic {
	return &PutAdminerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutAdminerLogic) PutAdminer(req *types.PutAdminerReq) (any, error) {
	in := &admin_pb.PutAdminerReq{
		Id:           req.Id,
		Name:         req.Name,
		Avatar:       req.Avatar,
		Access:       req.Access,
		Passport:     req.Passport,
		Password:     req.Password,
		Email:        req.Email,
		Status:       req.Status,
		IsSuperAdmin: req.IsSuperAdmin,
		LoginCount:   req.LoginCount,
		LastLogin:    req.LastLogin,
	}
	_, err := l.svcCtx.AdminRpc.PutAdminer(l.ctx, in)
	if err != nil {
		//若rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}

	return respx.DefaultStateResp(nil)
}
