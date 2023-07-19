package adminer

import (
	"context"
	"net/http"

	"mixbench/common/errorx"
	"mixbench/common/jwtx"
	"mixbench/common/respx"
	"mixbench/services/v1/admin/api/internal/svc"
	"mixbench/services/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CurrentInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewCurrentInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *CurrentInfoLogic {
	return &CurrentInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *CurrentInfoLogic) CurrentInfo() (any, error) {
	id, err := jwtx.GetIdByRequest(l.r, "")
	if err != nil {
		return nil, errorx.InternalError(err)
	}
	resp, err := l.svcCtx.AdminRpc.GetAdminer(l.ctx, &admin_pb.GetAdminerReq{
		Id: id,
	})
	if err != nil {
		return nil, err
	}
	resp.Adminer.Password = ""
	return respx.NewData(resp.Adminer), nil
}
