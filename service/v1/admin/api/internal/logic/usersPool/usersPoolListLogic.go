package usersPool

import (
	"context"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
	"mixbench/service/v1/admin/rpc/admin_pb"

	"mixbench/common/respx"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserPoolListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserPoolListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserPoolListLogic {
	return &UserPoolListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserPoolListLogic) UserPoolList(req *types.UserPoolListReq) (any, error) {
	resp, err := l.svcCtx.AdminRpc.GetUserPoolList(l.ctx, &admin_pb.UserPoolListReq{
		ListReq: &admin_pb.ListReq{
			PageSize: req.PageSize,
			Page:     req.Page,
			Keyword:  req.Keyword,
		},
		UserPool: &admin_pb.UserPool{
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
