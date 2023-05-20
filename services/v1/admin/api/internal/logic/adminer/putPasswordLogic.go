package adminer

import (
	"context"
	"errors"

	"mixbench/common/ctxdata"
	"mixbench/common/errorx"
	"mixbench/common/respx"
	"mixbench/services/v1/admin/api/internal/svc"
	"mixbench/services/v1/admin/api/internal/types"
	"mixbench/services/v1/admin/rpc/admin_pb"
	"mixbench/services/v1/captcha/rpc/captcha_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPutPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutPasswordLogic {
	return &PutPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutPasswordLogic) PutPassword(req *types.PutPasswordReq) (any, error) {
	//验证码校验
	resp1, err := l.svcCtx.CaptchaRpc.VerifyCaptcha(l.ctx, &captcha_pb.VerifyCaptchaReq{
		CaptchaId: req.CaptchaId,
		Solution:  req.Solution,
	})
	if err != nil {
		return nil, err
	}
	if !resp1.Success {
		return nil, errorx.ExternalError(nil, errors.New("验证码错误"))
	}
	id := ctxdata.GetUidFromCtx(l.ctx)
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
