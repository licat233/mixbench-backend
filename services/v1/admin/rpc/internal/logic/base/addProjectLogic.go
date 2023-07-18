package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/common/uniqueid"
	"mixbench/services/v1/admin/model"
	"mixbench/services/v1/admin/rpc/admin_pb"
	"mixbench/services/v1/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddProjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProjectLogic {
	return &AddProjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加项目
func (l *AddProjectLogic) AddProject(in *admin_pb.AddProjectReq) (*admin_pb.AddProjectResp, error) {
	data := &model.Project{
		Uuid:      uniqueid.NewUUID(), // 这里的uniqueid包，请自己定义
		Name:      in.Name,
		CompanyId: in.CompanyId,
		Status:    in.Status,
	}
	if _, err := l.svcCtx.ProjectModel.Insert(l.ctx, data); err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.AddProjectResp{}, nil
}
