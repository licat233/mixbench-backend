package user

import (
	"context"

	"mixbench/common/respx"
	"mixbench/services/v1/luckydraw/api/internal/svc"
	"mixbench/services/v1/luckydraw/api/internal/types"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

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

func (l *UserDetailsLogic) UserDetails(req *types.GetUserListReq) (any, error) {
	resp, err := l.svcCtx.LuckydrawRpc.GetUserDetails(l.ctx, &luckydraw_pb.GetUserDetailsReq{
		ListReq: &luckydraw_pb.ListReq{
			PageSize: req.PageSize,
			Page:     req.Page,
			Keyword:  req.Keyword,
		},
		User: &luckydraw_pb.User{
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

	data := []*types.UserDetail{}

	for _, item := range resp.UserDetails {
		data = append(data, PbUserDetailToApiUserDetail(item))
	}

	return respx.ListResp("", data, resp.Total, req.PageSize, req.Page), nil
}
