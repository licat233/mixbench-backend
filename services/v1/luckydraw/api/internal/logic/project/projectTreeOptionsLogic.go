package project

import (
	"context"

	"mixbench/common/respx"
	"mixbench/services/v1/luckydraw/api/internal/logic/dataconv"
	"mixbench/services/v1/luckydraw/api/internal/svc"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProjectTreeOptionsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProjectTreeOptionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProjectTreeOptionsLogic {
	return &ProjectTreeOptionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProjectTreeOptionsLogic) ProjectTreeOptions() (any, error) {
	resp, err := l.svcCtx.LuckydrawRpc.GetProjectTreeOptions(l.ctx, &luckydraw_pb.NilReq{})
	if err != nil {
		return nil, err
	}
	data := dataconv.PbTreeOptionsToApiTreeOptions(resp.TreeOptions)
	return respx.NewData(data), nil
}
