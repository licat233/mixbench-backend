package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/services/v1/luckydraw/rpc/internal/svc"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelProjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelProjectLogic {
	return &DelProjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除项目
func (l *DelProjectLogic) DelProject(in *luckydraw_pb.DelProjectReq) (*luckydraw_pb.DelProjectResp, error) {
	err := l.svcCtx.ProjectModel.SoftDelete(l.ctx, in.Id)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &luckydraw_pb.DelProjectResp{}, nil
}
