package lucker

import (
	"context"

	"mixbench/common/respx"
	"mixbench/service/v1/admin/api/internal/logic/common"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LuckerEnumListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLuckerEnumListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LuckerEnumListLogic {
	return &LuckerEnumListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LuckerEnumListLogic) LuckerEnumList() (any, error) {
	resp, err := l.svcCtx.AdminRpc.GetLuckerEnumList(l.ctx, &admin_pb.EnumListReq{})
	if err != nil {
		return nil, err
	}
	data := common.ConvertEnumList(resp.Data)
	return respx.DefaultSingleResp(data, err)
	// return common.ConvertEnumList(resp.Data), nil
}
