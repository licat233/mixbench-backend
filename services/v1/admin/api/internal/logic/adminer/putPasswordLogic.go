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
	"mixbench/services/v1/authentication/rpc/authentication_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewPutPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *PutPasswordLogic {
	return &PutPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *PutPasswordLogic) PutPassword(req *types.PutPasswordReq) (any, error) {
	//验证码校验
	resp1, err := l.svcCtx.AuthenticationRpc.VerifyCaptcha(l.ctx, &authentication_pb.VerifyCaptchaReq{
		CaptchaId: req.CaptchaId,
		Solution:  req.Solution,
	})
	if err != nil {
		return nil, err
	}
	if !resp1.Success {
		return nil, errorx.ExternalError(nil, errors.New("验证码错误"))
	}
	id, err := jwtx.GetIdByRequest(l.r, "")
	if err != nil {
		return nil, errorx.InternalError(err)
	}
	_, err = l.svcCtx.AdminRpc.PutPassword(l.ctx, &admin_pb.PutPasswordReq{
		AdminerId:   id,
		OldPassword: req.OldPassword,
		NewPassword: req.NewPassword,
		CaptchaId:   req.CaptchaId,
		Solution:    req.Solution,
	})
	if err != nil {
		return nil, err
	}
	return respx.StatusResp("", true), nil
}
