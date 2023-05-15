package logic

import (
	"context"
	"mixbench/services/v1/captcha/api/internal/svc"
	"mixbench/services/v1/captcha/api/internal/types"
	"mixbench/services/v1/captcha/rpc/captcha_pb"

	"mixbench/common/errorx"
	"mixbench/common/respx"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProcessCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProcessCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProcessCaptchaLogic {
	return &ProcessCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProcessCaptchaLogic) ProcessCaptcha(req *types.ProcessCaptchaReq) (any, error) {
	in := &captcha_pb.VerifyCaptchaReq{
		CaptchaId: req.CaptchaId,
		Solution:  req.Solution,
	}
	resp, err := l.svcCtx.CaptchaRpc.VerifyCaptcha(l.ctx, in)
	if err != nil {
		return nil, err
	}
	if resp.Success {
		return respx.StateResp("Captcha verification Success", true), nil
	}
	return nil, errorx.New("Captcha verification failed")
}
