package logic

import (
	"context"
	"mixbench/service/v1/admin/rpc/admin_pb"
	"mixbench/service/v1/admin/rpc/internal/svc"

	"mixbench/common/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserPoolListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserPoolListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserPoolListLogic {
	return &GetUserPoolListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserPoolListLogic) GetUserPoolList(in *admin_pb.UserPoolListReq) (*admin_pb.GetUserPoolListResp, error) {
	in.ListReq = newListReq(in.ListReq)
	list, total, err := l.svcCtx.UserPoolModel.FindList(l.ctx, in.ListReq.PageSize, in.ListReq.Page, pbUserPoolToMdUserPool(in.UserPool))
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.GetUserPoolListResp{
		Data:  mdUserPoolListToPbUserPoolList(list),
		Total: total,
	}, nil
}
