package logic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/service/v1/admin/rpc/admin_pb"
	"mixbench/service/v1/admin/rpc/internal/svc"
	"mixbench/service/v1/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAdminerEnumListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAdminerEnumListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdminerEnumListLogic {
	return &GetAdminerEnumListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAdminerEnumListLogic) GetAdminerEnumList(in *admin_pb.EnumListReq) (*admin_pb.EnumListResp, error) {
	list, _, err := l.svcCtx.AdminerModel.FindList(l.ctx, -1, -1, "", nil)
	if err != nil && err != model.ErrNotFound {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	data := []*admin_pb.Enum{}
	for _, item := range list {
		data = append(data, &admin_pb.Enum{
			Label: item.Name,
			Value: item.Id,
		})
	}
	return &admin_pb.EnumListResp{
		Data: data,
	}, nil
}
