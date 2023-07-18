package web

import (
	"context"

	"mixbench/common/respx"
	"mixbench/services/v1/admin/api/internal/svc"
	"mixbench/services/v1/admin/api/internal/types"
	"mixbench/services/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutWebLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPutWebLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutWebLogic {
	return &PutWebLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutWebLogic) PutWeb(req *types.PutWebReq) (any, error) {
	in := &admin_pb.PutWebReq{
		Web: &admin_pb.Web{
			Id:          req.Id,
			Uuid:        req.Uuid,
			ProjectId:   req.ProjectId,
			Url:         req.Url,
			Description: req.Description,
			Code:        req.Code,
		},
	}
	if _, err := l.svcCtx.AdminRpc.PutWeb(l.ctx, in); err != nil {
		//若rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}

	return respx.DefaultStatusResp(nil)
}
