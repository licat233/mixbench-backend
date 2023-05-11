package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/service/v1/luckydraw/model"
	"mixbench/service/v1/luckydraw/rpc/internal/svc"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

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
func (l *PutProjectLogic) PutProject(in *luckydraw_pb.PutProjectReq) (*luckydraw_pb.PutProjectResp, error) {
	data := &model.Project{
		Id:     in.Id,
		Uuid:   in.Uuid,
		Name:   in.Name,
		Status: in.Status,
	}

	if err := l.svcCtx.ProjectModel.Update(l.ctx, data); err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &luckydraw_pb.PutProjectResp{}, nil
}
