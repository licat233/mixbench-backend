package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/common/uniqueid"
	"mixbench/services/v1/admin/model"
	"mixbench/services/v1/admin/rpc/admin_pb"
	"mixbench/services/v1/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCompanyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCompanyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCompanyLogic {
	return &AddCompanyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加公司
func (l *AddCompanyLogic) AddCompany(in *admin_pb.AddCompanyReq) (*admin_pb.AddCompanyResp, error) {
	data := &model.Company{
		Uuid:   uniqueid.NewUUID(), // 这里的uniqueid包，请自己定义
		Name:   in.Name,
		Status: in.Status,
	}
	if _, err := l.svcCtx.CompanyModel.Insert(l.ctx, data); err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.AddCompanyResp{}, nil
}
