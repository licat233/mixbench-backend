/*
 * @Author: licat
 * @Date: 2023-02-01 11:56:03
 * @LastEditors: licat
 * @LastEditTime: 2023-02-01 15:42:41
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

type GetProjectActivitiesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProjectActivitiesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProjectActivitiesLogic {
	return &GetProjectActivitiesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 某项目的活动
func (l *GetProjectActivitiesLogic) GetProjectActivities(in *luckydraw_pb.GetProjectActivitiesReq) (*luckydraw_pb.GetProjectActivitiesResp, error) {
	list, _, err := l.svcCtx.ActivityModel.FindList(l.ctx, -1, -1, "", &model.Activity{
		ProjectId: in.ProjectId,
	})
	if err != nil && err != model.ErrNotFound {
		return nil, errorx.IntRpcErr(err)
	}
	return &luckydraw_pb.GetProjectActivitiesResp{
		Activities: tools.MdActivitiesToPbActivities(list),
	}, nil
}
