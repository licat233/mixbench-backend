package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/services/v1/admin/rpc/admin_pb"
	"mixbench/services/v1/admin/rpc/internal/logic/dataconv"
	"mixbench/services/v1/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutProjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPutProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutProjectLogic {
	return &PutProjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新项目
func (l *PutProjectLogic) PutProject(in *admin_pb.PutProjectReq) (*admin_pb.PutProjectResp, error) {
	data := dataconv.PbProjectToMdProject(in.Project)
	if err := l.svcCtx.ProjectModel.Update(l.ctx, data); err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.PutProjectResp{}, nil
}
