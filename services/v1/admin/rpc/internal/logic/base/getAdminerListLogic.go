package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/services/v1/admin/rpc/admin_pb"
	"mixbench/services/v1/admin/rpc/internal/logic/dataconv"
	"mixbench/services/v1/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAdminerListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAdminerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdminerListLogic {
	return &GetAdminerListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取管理员列表
func (l *GetAdminerListLogic) GetAdminerList(in *admin_pb.GetAdminerListReq) (*admin_pb.GetAdminerListResp, error) {
	pageSize, page, keyword := dataconv.ListReqParams(in.ListReq)
	list, total, err := l.svcCtx.AdminerModel.FindList(l.ctx, pageSize, page, keyword, dataconv.PbAdminerToMdAdminer(in.Adminer))
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.GetAdminerListResp{
		Adminers: dataconv.MdAdminers2PbAdminers(list),
		Total:    total,
	}, nil
}
