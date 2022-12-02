package activity

import (
	"context"

	"mixbench/common/respx"
	"mixbench/service/v1/admin/api/internal/logic/common"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActivityPatternEnumListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewActivityPatternEnumListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActivityPatternEnumListLogic {
	return &ActivityPatternEnumListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ActivityPatternEnumListLogic) ActivityPatternEnumList() (any, error) {
	resp, err := l.svcCtx.AdminRpc.ActivityPatternEnumList(l.ctx, &admin_pb.NilReq{})
	if err != nil {
		return nil, err
	}
	data := common.ConvertEnumList(resp.Data)
	return respx.DefaultSingleResp(data, err)
	// return common.ConvertEnumList(resp.Data), nil
}
