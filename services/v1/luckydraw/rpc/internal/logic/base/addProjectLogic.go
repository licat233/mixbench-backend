package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/common/uniqueid"
	"mixbench/services/v1/luckydraw/model"
	"mixbench/services/v1/luckydraw/rpc/internal/svc"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

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
func (l *AddProjectLogic) AddProject(in *luckydraw_pb.AddProjectReq) (*luckydraw_pb.AddProjectResp, error) {
	data := &model.Project{
		Uuid:   uniqueid.NewUUID(),
		Name:   in.Name,
		Status: in.Status,
	}
	_, err := l.svcCtx.ProjectModel.Insert(l.ctx, data)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &luckydraw_pb.AddProjectResp{}, nil
}
