package logic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/service/v1/admin/rpc/admin_pb"
	"mixbench/service/v1/admin/rpc/internal/svc"
	"mixbench/service/v1/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetActivityEnumListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetActivityEnumListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetActivityEnumListLogic {
	return &GetActivityEnumListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetActivityEnumListLogic) GetActivityEnumList(in *admin_pb.GetActivityEnumListReq) (*admin_pb.EnumListResp, error) {
	var list = []*model.Activity{}
	var err error
	if in.ProjectId == -1 {
		list, _, err = l.svcCtx.ActivityModel.FindList(l.ctx, -1, -1, "", nil)
		if err != nil {
			logx.Error(err)
			return nil, errorx.IntRpcErr(err)
		}
	} else {
		list, err = l.svcCtx.ActivityModel.FindsByProjectId(l.ctx, in.ProjectId)
		if err != nil {
			logx.Error(err)
			return nil, errorx.IntRpcErr(err)
		}
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
