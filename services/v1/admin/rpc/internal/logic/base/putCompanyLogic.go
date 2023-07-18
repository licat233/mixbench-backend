package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/services/v1/admin/rpc/admin_pb"
	"mixbench/services/v1/admin/rpc/internal/logic/dataconv"
	"mixbench/services/v1/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutCompanyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPutCompanyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutCompanyLogic {
	return &PutCompanyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新公司
func (l *PutCompanyLogic) PutCompany(in *admin_pb.PutCompanyReq) (*admin_pb.PutCompanyResp, error) {
	data := dataconv.PbCompanyToMdCompany(in.Company)
	if err := l.svcCtx.CompanyModel.Update(l.ctx, data); err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.PutCompanyResp{}, nil
}
