package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/services/v1/authentication/rpc/authentication_pb"
	"mixbench/services/v1/authentication/rpc/internal/logic/dataconv"
	"mixbench/services/v1/authentication/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetIpBlacklistListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetIpBlacklistListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetIpBlacklistListLogic {
	return &GetIpBlacklistListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取IP黑名单列表
func (l *GetIpBlacklistListLogic) GetIpBlacklistList(in *authentication_pb.GetIpBlacklistListReq) (*authentication_pb.GetIpBlacklistListResp, error) {
	pageSize, page, keyword := dataconv.ListReqParams(in.ListReq)
	list, total, err := l.svcCtx.IpBlacklistModel.FindList(l.ctx, pageSize, page, keyword, dataconv.PbIpBlacklistToMdIpBlacklist(in.IpBlacklist))
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &authentication_pb.GetIpBlacklistListResp{
		IpBlacklists: dataconv.MdIpBlacklists2PbIpBlacklists(list),
		Total:        total,
	}, nil
}
