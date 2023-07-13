package web

import (
	"context"

	"mixbench/common/respx"
	"mixbench/services/v1/admin/api/internal/svc"
	"mixbench/services/v1/admin/api/internal/types"
	"mixbench/services/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelWebLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelWebLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelWebLogic {
	return &DelWebLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelWebLogic) DelWeb(req *types.DelWebReq) (any, error) {
	_, err := l.svcCtx.AdminRpc.DelWeb(l.ctx, &admin_pb.DelWebReq{
		Id: req.Id,
	})
	if err != nil {
		//若rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}

	return respx.DefaultStatusResp(nil)
}
