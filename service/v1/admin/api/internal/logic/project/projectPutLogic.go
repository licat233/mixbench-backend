package project

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/common/respx"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
	"mixbench/service/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProjectPutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProjectPutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProjectPutLogic {
	return &ProjectPutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProjectPutLogic) ProjectPut(req *types.ProjectPutReq) (any, error) {
	if req.Name == "" {
		return nil, errorx.ErrInvalidParam
	}
	_, err := l.svcCtx.AdminRpc.PutProject(l.ctx, &admin_pb.Project{
		Id:     req.Id,
		Uuid:   req.Uuid,
		Name:   req.Name,
		Status: req.Status,
	})
	if err != nil {
		return nil, err
	}
	return respx.DefaultStateResp(err)
}
