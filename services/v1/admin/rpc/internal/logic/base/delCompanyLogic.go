package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/services/v1/admin/rpc/admin_pb"
	"mixbench/services/v1/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelCompanyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelCompanyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelCompanyLogic {
	return &DelCompanyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除公司
func (l *DelCompanyLogic) DelCompany(in *admin_pb.DelCompanyReq) (*admin_pb.DelCompanyResp, error) {
	if err := l.svcCtx.CompanyModel.SoftDelete(l.ctx, in.Id); err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.DelCompanyResp{}, nil
}
