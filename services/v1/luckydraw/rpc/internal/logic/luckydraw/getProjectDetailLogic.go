/*
 * @Author: licat
 * @Date: 2023-01-31 17:18:24
 * @LastEditors: licat
 * @LastEditTime: 2023-02-01 15:43:35
 * @Description: licat233@gmail.com
 */
package luckydrawlogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/services/v1/luckydraw/model"
	"mixbench/services/v1/luckydraw/rpc/internal/logic/tools"
	"mixbench/services/v1/luckydraw/rpc/internal/svc"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProjectDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProjectDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProjectDetailLogic {
	return &GetProjectDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取项目详细
func (l *GetProjectDetailLogic) GetProjectDetail(in *luckydraw_pb.GetProjectDetailReq) (*luckydraw_pb.GetProjectDetailResp, error) {
	project, err := l.svcCtx.ProjectModel.FindOne(l.ctx, in.ProjectId)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	activities, _, err := l.svcCtx.ActivityModel.FindList(l.ctx, -1, -1, "", &model.Activity{
		ProjectId: in.ProjectId,
	})
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &luckydraw_pb.GetProjectDetailResp{
		ProjectDetail: &luckydraw_pb.ProjectDetail{
			Project:    tools.MdProjectToPbProject(project),
			Activities: tools.MdActivitiesToPbActivities(activities),
		},
	}, nil
}
