package adminer

import (
	"context"

	"mixbench/common/respx"
	"mixbench/service/v1/admin/api/internal/logic/dataconv"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
	"mixbench/service/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAdminerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAdminerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdminerLogic {
	return &GetAdminerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAdminerLogic) GetAdminer(req *types.GetAdminerReq) (any, error) {
	rpcResp, err := l.svcCtx.AdminRpc.GetAdminer(l.ctx, &admin_pb.GetAdminerReq{
		Id: req.Id,
	})
	if err != nil {
		//若rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}
	pbAdminer := rpcResp.Adminer
	data := dataconv.PbAdminerToApiAdminer(pbAdminer)

	return respx.DefaultSingleResp(data, nil)
}
