package logic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/service/v1/admin/rpc/admin_pb"
	"mixbench/service/v1/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProjectTreeOptionsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProjectTreeOptionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProjectTreeOptionsLogic {
	return &GetProjectTreeOptionsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProjectTreeOptionsLogic) GetProjectTreeOptions(in *admin_pb.EnumListReq) (*admin_pb.ProjectTreeOptionsResp, error) {
	list, _, err := l.svcCtx.ProjectModel.FindList(l.ctx, -1, -1, "", nil)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	data := []*admin_pb.TreeOptions{}
	for _, project := range list {
		activities, err := l.svcCtx.ActivityModel.FindsByProjectId(l.ctx, project.Id)
		if err != nil {
			l.Logger.Error(err)
			return nil, errorx.IntRpcErr(err)
		}
		activityOptions := []*admin_pb.TreeOptions{}
		for _, activity := range activities {
			activityOptions = append(activityOptions, &admin_pb.TreeOptions{
				Title:    activity.Name,
				Value:    activity.Id,
				Children: nil,
			})
		}
		data = append(data, &admin_pb.TreeOptions{
			Title:    project.Name,
			Value:    project.Id,
			Children: activityOptions,
		})
	}

	return &admin_pb.ProjectTreeOptionsResp{
		Data: data,
	}, nil
}
