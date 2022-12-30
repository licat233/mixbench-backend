package logic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/service/v1/admin/rpc/admin_pb"
	"mixbench/service/v1/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProjectDetailsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProjectDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProjectDetailsLogic {
	return &GetProjectDetailsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProjectDetailsLogic) GetProjectDetails(in *admin_pb.ProjectListReq) (*admin_pb.GetProjectDetailsListResp, error) {
	in.ListReq = newListReq(in.ListReq)
	list, total, err := l.svcCtx.ProjectModel.FindList(l.ctx, in.ListReq.PageSize, in.ListReq.Page, in.ListReq.Keyword, pbProjectToMdProject(in.Project))
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	data := []*admin_pb.ProjectDetails{}
	for _, project := range list {
		activities, err := l.svcCtx.ActivityModel.FindsByProjectId(l.ctx, project.Id)
		if err != nil {
			l.Logger.Error(err)
			return nil, errorx.IntRpcErr(err)
		}
		data = append(data, &admin_pb.ProjectDetails{
			Project:    mdProjectToPbProject(project),
			Activities: mdActivityListToPbActivityList(activities),
		})
	}
	return &admin_pb.GetProjectDetailsListResp{
		Data:  data,
		Total: total,
	}, nil
}
