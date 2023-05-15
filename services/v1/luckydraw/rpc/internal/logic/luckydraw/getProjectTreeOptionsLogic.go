/*
 * @Author: licat
 * @Date: 2023-01-31 17:18:24
 * @LastEditors: licat
 * @LastEditTime: 2023-02-20 17:45:42
 * @Description: licat233@gmail.com
 */
package luckydrawlogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/services/v1/luckydraw/model"
	"mixbench/services/v1/luckydraw/rpc/internal/svc"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

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

// 项目选项数据
func (l *GetProjectTreeOptionsLogic) GetProjectTreeOptions(in *luckydraw_pb.NilReq) (*luckydraw_pb.TreeOptions, error) {
	list, _, err := l.svcCtx.ProjectModel.FindList(l.ctx, -1, -1, "", nil)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	data := []*luckydraw_pb.TreeOption{}
	for _, project := range list {
		activities, _, err := l.svcCtx.ActivityModel.FindList(l.ctx, -1, -1, "", &model.Activity{
			ProjectId: project.Id,
		})
		if err != nil {
			l.Logger.Error(err)
			return nil, errorx.IntRpcErr(err)
		}
		activityOptions := []*luckydraw_pb.TreeOption{}
		for _, activity := range activities {
			activityOptions = append(activityOptions, &luckydraw_pb.TreeOption{
				Title:    activity.Name,
				Value:    activity.Id,
				Children: nil,
			})
		}
		data = append(data, &luckydraw_pb.TreeOption{
			Title:    project.Name,
			Value:    project.Id,
			Children: activityOptions,
		})
	}

	return &luckydraw_pb.TreeOptions{
		TreeOptions: data,
	}, nil
}
