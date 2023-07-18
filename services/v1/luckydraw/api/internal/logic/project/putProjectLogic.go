package project

import (
	"context"

	"mixbench/common/respx"
	"mixbench/services/v1/luckydraw/api/internal/svc"
	"mixbench/services/v1/luckydraw/api/internal/types"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPutProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutProjectLogic {
	return &PutProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutProjectLogic) PutProject(req *types.PutProjectReq) (any, error) {
	in := &luckydraw_pb.PutProjectReq{
		Project: &luckydraw_pb.Project{
			Id:     req.Id,
			Uuid:   req.Uuid,
			Name:   req.Name,
			Status: req.Status,
		},
	}
	if _, err := l.svcCtx.LuckydrawRpc.PutProject(l.ctx, in); err != nil {
		//若rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}

	return respx.DefaultStatusResp(nil)
}
