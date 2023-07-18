package company

import (
	"context"

	"mixbench/common/respx"
	"mixbench/services/v1/admin/api/internal/svc"
	"mixbench/services/v1/admin/api/internal/types"
	"mixbench/services/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCompanyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddCompanyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCompanyLogic {
	return &AddCompanyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddCompanyLogic) AddCompany(req *types.AddCompanyReq) (any, error) {
	in := &admin_pb.AddCompanyReq{
		Name:   req.Name,
		Status: req.Status,
	}
	if _, err := l.svcCtx.AdminRpc.AddCompany(l.ctx, in); err != nil {
		//若rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}

	return respx.DefaultStatusResp(nil)
}
