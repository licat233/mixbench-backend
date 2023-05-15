package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/services/v1/luckydraw/rpc/internal/logic/dataconv"
	"mixbench/services/v1/luckydraw/rpc/internal/svc"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProjectListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProjectListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProjectListLogic {
	return &GetProjectListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取项目列表
func (l *GetProjectListLogic) GetProjectList(in *luckydraw_pb.GetProjectListReq) (*luckydraw_pb.GetProjectListResp, error) {
	pageSize, page, keyword := dataconv.ListReqParams(in.ListReq)
	list, total, err := l.svcCtx.ProjectModel.FindList(l.ctx, pageSize, page, keyword, dataconv.PbProjectToMdProject(in.Project))
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &luckydraw_pb.GetProjectListResp{
		Projects: dataconv.MdProjects2PbProjects(list),
		Total:    total,
	}, nil
}
