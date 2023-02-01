/*
 * @Author: licat
 * @Date: 2023-01-31 17:18:24
 * @LastEditors: licat
 * @LastEditTime: 2023-01-31 17:36:08
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

type BaseGetProjectListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBaseGetProjectListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BaseGetProjectListLogic {
	return &BaseGetProjectListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BaseGetProjectListLogic) BaseGetProjectList(in *luckydraw_pb.GetProjectListReq) (*luckydraw_pb.GetProjectListResp, error) {
	in.ListReq = tools.NewListReq(in.ListReq)
	list, total, err := l.svcCtx.ProjectModel.FindList(l.ctx, in.ListReq.PageSize, in.ListReq.Page, in.ListReq.Keyword, tools.PbProjectToMdProject(in.Project))
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &luckydraw_pb.GetProjectListResp{
		Projects: tools.MdProjectsToPbProjects(list),
		Total:    total,
	}, nil
}
