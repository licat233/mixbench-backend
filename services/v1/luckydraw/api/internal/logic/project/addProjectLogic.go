package project

import (
	"context"

	"mixbench/common/respx"
	"mixbench/services/v1/luckydraw/api/internal/svc"
	"mixbench/services/v1/luckydraw/api/internal/types"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProjectLogic {
	return &AddProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddProjectLogic) AddProject(req *types.AddProjectReq) (any, error) {
	in := &luckydraw_pb.AddProjectReq{
		Name:   req.Name,
		Status: req.Status,
	}
	_, err := l.svcCtx.LuckydrawRpc.AddProject(l.ctx, in)
	if err != nil {
		//若rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}

	return respx.DefaultStatusResp(nil)
}
