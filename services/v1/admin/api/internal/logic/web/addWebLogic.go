package web

import (
	"context"

	"mixbench/common/respx"
	"mixbench/services/v1/admin/api/internal/svc"
	"mixbench/services/v1/admin/api/internal/types"
	"mixbench/services/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddWebLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddWebLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddWebLogic {
	return &AddWebLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddWebLogic) AddWeb(req *types.AddWebReq) (any, error) {
	in := &admin_pb.AddWebReq{
		ProjectId:   req.ProjectId,
		Url:         req.Url,
		Host:        req.Host,
		Description: req.Description,
		Code:        req.Code,
	}
	if _, err := l.svcCtx.AdminRpc.AddWeb(l.ctx, in); err != nil {
		//若rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}

	return respx.DefaultStatusResp(nil)
}
