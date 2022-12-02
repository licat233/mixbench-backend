package project

import (
	"context"

	"mixbench/common/respx"
	"mixbench/service/v1/admin/api/internal/logic/common"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProjectEnumListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProjectEnumListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProjectEnumListLogic {
	return &ProjectEnumListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProjectEnumListLogic) ProjectEnumList() (any, error) {
	resp, err := l.svcCtx.AdminRpc.GetProjectEnumList(l.ctx, &admin_pb.EnumListReq{})
	if err != nil {
		return nil, err
	}
	data := common.ConvertEnumList(resp.Data)
	return respx.DefaultSingleResp(data, err)
	// return common.ConvertEnumList(resp.Data), nil
}
