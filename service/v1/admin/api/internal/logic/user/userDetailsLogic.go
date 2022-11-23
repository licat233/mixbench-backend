package user

import (
	"context"

	"mixbench/common/respx"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
	"mixbench/service/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDetailsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailsLogic {
	return &UserDetailsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDetailsLogic) UserDetails(req *types.UserListReq) (any, error) {
	resp, err := l.svcCtx.AdminRpc.GetUserDetails(l.ctx, &admin_pb.UserListReq{
		ListReq: &admin_pb.ListReq{
			PageSize: req.PageSize,
			Page:     req.Page,
			Keyword:  req.Keyword,
		},
		User: &admin_pb.User{
			Passport: req.Passport,
			Email:    req.Email,
			Phone:    req.Phone,
			Name:     req.Name,
			Address:  req.Address,
			Score:    req.Score,
			Status:   req.Status,
		},
	})
	if err != nil {
		return nil, err
	}

	return respx.ListResp("", resp.Data, resp.Total, req.PageSize, req.Page), nil
}
