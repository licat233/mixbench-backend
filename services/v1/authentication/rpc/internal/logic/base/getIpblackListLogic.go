package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/services/v1/authentication/rpc/authentication_pb"
	"mixbench/services/v1/authentication/rpc/internal/logic/dataconv"
	"mixbench/services/v1/authentication/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetIpblackListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetIpblackListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetIpblackListLogic {
	return &GetIpblackListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取IP黑名单列表
func (l *GetIpblackListLogic) GetIpblackList(in *authentication_pb.GetIpblackListReq) (*authentication_pb.GetIpblackListResp, error) {
	pageSize, page, keyword := dataconv.ListReqParams(in.ListReq)
	list, total, err := l.svcCtx.IpblackModel.FindList(l.ctx, pageSize, page, keyword, dataconv.PbIpblackToMdIpblack(in.Ipblack))
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &authentication_pb.GetIpblackListResp{
		Ipblacks: dataconv.MdIpblacks2PbIpblacks(list),
		Total:    total,
	}, nil
}
