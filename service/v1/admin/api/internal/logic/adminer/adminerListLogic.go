package adminer

import (
	"context"
	"errors"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
	"mixbench/service/v1/admin/rpc/admin_pb"

	"mixbench/common/errorx"
	"mixbench/common/respx"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminerListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminerListLogic {
	return &AdminerListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminerListLogic) AdminerList(req *types.AdminerListReq) (any, error) {
	ok, err := SuperAdminerIdentify(l.ctx, l.svcCtx)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errorx.AccessError(nil, errors.New("no request permission"))
	}
	resp, err := l.svcCtx.AdminRpc.GetAdminerList(l.ctx, &admin_pb.AdminerListReq{
		ListReq: &admin_pb.ListReq{
			PageSize: req.PageSize,
			Page:     req.Page,
			Keyword:  req.Keyword,
		},
		Adminer: &admin_pb.Adminer{
			Name:         req.Name,
			Passport:     req.Passport,
			Password:     req.Password,
			Email:        req.Email,
			Status:       req.Status,
			IsSuperAdmin: req.IsSuperAdmin,
			LoginCount:   req.LoginCount,
			LastLogin:    req.LastLogin,
		},
	})
	if err != nil {
		return nil, err
	}
	return respx.DefaultListResp(resp.Data, resp.Total, req.PageSize, req.Page, err)
}
