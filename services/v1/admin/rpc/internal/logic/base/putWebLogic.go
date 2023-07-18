package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/services/v1/admin/rpc/admin_pb"
	"mixbench/services/v1/admin/rpc/internal/logic/dataconv"
	"mixbench/services/v1/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutWebLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPutWebLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutWebLogic {
	return &PutWebLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新网站
func (l *PutWebLogic) PutWeb(in *admin_pb.PutWebReq) (*admin_pb.PutWebResp, error) {
	data := dataconv.PbWebToMdWeb(in.Web)
	if err := l.svcCtx.WebModel.Update(l.ctx, data); err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.PutWebResp{}, nil
}
