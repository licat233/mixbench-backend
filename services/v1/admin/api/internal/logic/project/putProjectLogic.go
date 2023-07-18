package project

import (
	"context"

	"mixbench/common/respx"
	"mixbench/services/v1/admin/api/internal/svc"
	"mixbench/services/v1/admin/api/internal/types"
	"mixbench/services/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPutProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutProjectLogic {
	return &PutProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutProjectLogic) PutProject(req *types.PutProjectReq) (any, error) {
	in := &admin_pb.PutProjectReq{
		Project: &admin_pb.Project{
			Id:        req.Id,
			Uuid:      req.Uuid,
			Name:      req.Name,
			CompanyId: req.CompanyId,
			Status:    req.Status,
		},
	}
	if _, err := l.svcCtx.AdminRpc.PutProject(l.ctx, in); err != nil {
		//若rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}

	return respx.DefaultStatusResp(nil)
}
