package adminer

import (
	"context"
	"errors"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
	"mixbench/service/v1/admin/rpc/admin_pb"

	"mixbench/common/errorx"
	"mixbench/common/respx"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminerDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminerDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminerDelLogic {
	return &AdminerDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminerDelLogic) AdminerDel(req *types.AdminerDelReq) (any, error) {
	ok, err := SuperAdminerIdentify(l.ctx, l.svcCtx)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errorx.AccessError(nil, errors.New("no request permission"))
	}
	_, err = l.svcCtx.AdminRpc.DelAdminer(l.ctx, &admin_pb.ItemReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return respx.DefaultStateResp(err)
}
