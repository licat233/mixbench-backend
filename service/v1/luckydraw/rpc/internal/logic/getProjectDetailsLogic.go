/*
 * @Author: licat
 * @Date: 2023-02-01 14:59:03
 * @LastEditors: licat
 * @LastEditTime: 2023-02-01 15:43:10
 * @Description: licat233@gmail.com
 */
package logic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/service/v1/luckydraw/rpc/internal/logic/tools"
	"mixbench/service/v1/luckydraw/rpc/internal/svc"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

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

// 获取项目详细列表
func (l *GetProjectDetailsLogic) GetProjectDetails(in *luckydraw_pb.GetProjectDetailsReq) (*luckydraw_pb.GetProjectDetailsResp, error) {
	in.ListReq = tools.NewListReq(in.ListReq)
	list, total, err := l.svcCtx.ProjectModel.FindList(l.ctx, in.ListReq.PageSize, in.ListReq.Page, in.ListReq.Keyword, tools.PbProjectToMdProject(in.Project))
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	data := []*luckydraw_pb.ProjectDetail{}
	for _, project := range list {
		activities, err := l.svcCtx.ActivityModel.FindsByProjectId(l.ctx, project.Id)
		if err != nil {
			l.Logger.Error(err)
			return nil, errorx.IntRpcErr(err)
		}
		data = append(data, &luckydraw_pb.ProjectDetail{
			Project:    tools.MdProjectToPbProject(project),
			Activities: tools.MdActivitiesToPbActivities(activities),
		})
	}
	return &luckydraw_pb.GetProjectDetailsResp{
		ProjectDetails: data,
		Total:          total,
	}, nil
}
