package logic

import (
	"context"
	"mixbench/service/v1/admin/rpc/admin_pb"
	"mixbench/service/v1/admin/rpc/internal/svc"

	"mixbench/common/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUsersPoolListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUsersPoolListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUsersPoolListLogic {
	return &GetUsersPoolListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUsersPoolListLogic) GetUsersPoolList(in *admin_pb.UsersPoolListReq) (*admin_pb.GetUsersPoolListResp, error) {
	in.ListReq = newListReq(in.ListReq)
	list, total, err := l.svcCtx.UsersPoolModel.FindList(l.ctx, in.ListReq.PageSize, in.ListReq.Page, pbUsersPoolToMdUsersPool(in.UsersPool))
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.GetUsersPoolListResp{
		Data:  mdUsersPoolListToPbUsersPoolList(list),
		Total: total,
	}, nil
}
