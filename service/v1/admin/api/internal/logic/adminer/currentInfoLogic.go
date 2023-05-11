package adminer

import (
	"context"

	"mixbench/common/ctxdata"
	"mixbench/common/respx"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CurrentInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCurrentInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CurrentInfoLogic {
	return &CurrentInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CurrentInfoLogic) CurrentInfo() (any, error) {
	id := ctxdata.GetUidFromCtx(l.ctx)
	resp, err := l.svcCtx.AdminRpc.GetAdminer(l.ctx, &admin_pb.GetAdminerReq{
		Id: id,
	})
	if err != nil {
		return nil, err
	}
	resp.Adminer.Password = ""
	return respx.NewData(resp.Adminer), nil
}
