package activity

import (
	"context"

	"mixbench/common/respx"
	"mixbench/service/v1/admin/api/internal/logic/common"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
	"mixbench/service/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActivityEnumListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewActivityEnumListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActivityEnumListLogic {
	return &ActivityEnumListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ActivityEnumListLogic) ActivityEnumList(req *types.EnumListReq) (any, error) {
	resp, err := l.svcCtx.AdminRpc.GetActivityEnumList(l.ctx, &admin_pb.GetActivityEnumListReq{
		ProjectId: req.ProjectId,
	})
	data := common.ConvertEnumList(resp.Data)
	return respx.DefaultSingleResp(data, err)
	// return common.ConvertEnumList(resp.Data), nil
}
