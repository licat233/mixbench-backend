package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/services/v1/admin/model"
	"mixbench/services/v1/admin/rpc/admin_pb"
	"mixbench/services/v1/admin/rpc/internal/logic/dataconv"
	"mixbench/services/v1/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetWebLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetWebLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWebLogic {
	return &GetWebLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取网站
func (l *GetWebLogic) GetWeb(in *admin_pb.GetWebReq) (*admin_pb.GetWebResp, error) {
	res, err := l.svcCtx.WebModel.FindById(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.GetWebResp{
		Web: dataconv.MdWebToPbWeb(res),
	}, nil
}
