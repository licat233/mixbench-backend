package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/services/v1/admin/rpc/admin_pb"
	"mixbench/services/v1/admin/rpc/internal/logic/dataconv"
	"mixbench/services/v1/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetWebListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetWebListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWebListLogic {
	return &GetWebListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取网站列表
func (l *GetWebListLogic) GetWebList(in *admin_pb.GetWebListReq) (*admin_pb.GetWebListResp, error) {
	pageSize, page, keyword := dataconv.ListReqParams(in.ListReq)
	list, total, err := l.svcCtx.WebModel.FindList(l.ctx, pageSize, page, keyword, dataconv.PbWebToMdWeb(in.Web))
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.GetWebListResp{
		Webs:  dataconv.MdWebs2PbWebs(list),
		Total: total,
	}, nil
}
