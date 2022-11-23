package project

import (
	"context"

	"mixbench/service/v1/admin/api/internal/logic/common"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
	"mixbench/service/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type EnumListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEnumListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EnumListLogic {
	return &EnumListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EnumListLogic) EnumList() ([]types.Enum, error) {
	resp, err := l.svcCtx.AdminRpc.GetProjectEnumList(l.ctx, &admin_pb.EnumListReq{})
	if err != nil {
		return nil, err
	}
	return common.ConvertEnumList(resp.Data), nil
}
