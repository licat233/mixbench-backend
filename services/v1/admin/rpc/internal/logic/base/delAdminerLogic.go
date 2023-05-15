package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/services/v1/admin/rpc/admin_pb"
	"mixbench/services/v1/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelAdminerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelAdminerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelAdminerLogic {
	return &DelAdminerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除管理员
func (l *DelAdminerLogic) DelAdminer(in *admin_pb.DelAdminerReq) (*admin_pb.DelAdminerResp, error) {
	err := l.svcCtx.AdminerModel.Delete(l.ctx, in.Id)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.DelAdminerResp{}, nil
}
