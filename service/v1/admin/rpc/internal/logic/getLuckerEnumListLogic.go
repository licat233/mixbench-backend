package logic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/service/v1/admin/rpc/admin_pb"
	"mixbench/service/v1/admin/rpc/internal/svc"
	"mixbench/service/v1/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLuckerEnumListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLuckerEnumListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLuckerEnumListLogic {
	return &GetLuckerEnumListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLuckerEnumListLogic) GetLuckerEnumList(in *admin_pb.EnumListReq) (*admin_pb.EnumListResp, error) {
	list, _, err := l.svcCtx.LuckerModel.FindList(l.ctx, -1, -1, "", nil)
	if err != nil && err != model.ErrNotFound {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	data := []*admin_pb.Enum{}
	for _, item := range list {
		data = append(data, &admin_pb.Enum{
			Label: item.Uuid,
			Value: item.Id,
		})
	}
	return &admin_pb.EnumListResp{
		Data: data,
	}, nil
}
