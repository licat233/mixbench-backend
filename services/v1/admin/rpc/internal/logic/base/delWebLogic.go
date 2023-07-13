package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/services/v1/admin/rpc/admin_pb"
	"mixbench/services/v1/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelWebLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelWebLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelWebLogic {
	return &DelWebLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除网站
func (l *DelWebLogic) DelWeb(in *admin_pb.DelWebReq) (*admin_pb.DelWebResp, error) {
	if err := l.svcCtx.WebModel.SoftDelete(l.ctx, in.Id); err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.DelWebResp{}, nil
}
