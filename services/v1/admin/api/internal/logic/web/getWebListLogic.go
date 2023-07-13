package web

import (
	"context"

	"mixbench/common/respx"
	"mixbench/services/v1/admin/api/internal/logic/dataconv"
	"mixbench/services/v1/admin/api/internal/svc"
	"mixbench/services/v1/admin/api/internal/types"
	"mixbench/services/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetWebListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetWebListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWebListLogic {
	return &GetWebListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetWebListLogic) GetWebList(req *types.GetWebListReq) (any, error) {
	in := &admin_pb.GetWebListReq{
		ListReq: &admin_pb.ListReq{
			PageSize: req.PageSize,
			Page:     req.Page,
			Keyword:  req.Keyword,
		},
		Web: &admin_pb.Web{
			Id:          req.Id,
			Uuid:        req.Uuid,
			ProjectId:   req.ProjectId,
			Url:         req.Url,
			Host:        req.Host,
			Description: req.Description,
			Code:        req.Code,
		},
	}
	rpcResp, err := l.svcCtx.AdminRpc.GetWebList(l.ctx, in)
	if err != nil {
		//rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}
	pbList := rpcResp.Webs
	data := dataconv.PbWebs2ApiWebs(pbList)

	return respx.DefaultListResp(data, rpcResp.Total, req.PageSize, req.Page, nil)
}
