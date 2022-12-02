package usersPool

import (
	"context"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
	"mixbench/service/v1/admin/rpc/admin_pb"

	"mixbench/common/respx"

	"github.com/zeromicro/go-zero/core/logx"
)

type UsersPoolListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUsersPoolListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UsersPoolListLogic {
	return &UsersPoolListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UsersPoolListLogic) UsersPoolList(req *types.UsersPoolListReq) (any, error) {
	resp, err := l.svcCtx.AdminRpc.GetUsersPoolList(l.ctx, &admin_pb.UsersPoolListReq{
		ListReq: &admin_pb.ListReq{
			PageSize: req.PageSize,
			Page:     req.Page,
			Keyword:  req.Keyword,
		},
		UsersPool: &admin_pb.UsersPool{
			ActivityId: req.ActivityId,
			UserId:     req.UserId,
			Status:     req.Status,
			Score:      req.Score,
		},
	})
	if err != nil {
		return nil, err
	}
	return respx.DefaultListResp(resp.Data, resp.Total, req.PageSize, req.Page, err)
}
