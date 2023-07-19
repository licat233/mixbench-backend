package adminer

import (
	"context"
	"errors"
	"net/http"

	"mixbench/common/errorx"
	"mixbench/common/jwtx"
	"mixbench/common/respx"
	"mixbench/services/v1/admin/api/internal/svc"
	"mixbench/services/v1/admin/api/internal/types"
	"mixbench/services/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutCurrentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewPutCurrentLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *PutCurrentLogic {
	return &PutCurrentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *PutCurrentLogic) PutCurrent(req *types.PutCurrentReq) (any, error) {
	if len(req.Avatar) > 250 {
		return nil, errorx.ExternalError(errors.New("the value length of the 'avatar' field is too long"))
	}
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

	adminer := resp.Adminer
	adminer.Name = req.Name
	adminer.Avatar = req.Avatar
	adminer.Passport = req.Passport
	adminer.Email = req.Email
	_, err = l.svcCtx.AdminRpc.PutAdminer(l.ctx, &admin_pb.PutAdminerReq{
		Adminer: adminer,
	})
	if err != nil {
		return nil, err
	}
	return respx.DefaultStatusResp(err)
}
