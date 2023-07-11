package logic

import (
	"context"

	"mixbench/services/v1/authentication/rpc/authentication_pb"
	"mixbench/services/v1/authentication/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyCaptchaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVerifyCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyCaptchaLogic {
	return &VerifyCaptchaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 校验验证码
func (l *VerifyCaptchaLogic) VerifyCaptcha(in *authentication_pb.VerifyCaptchaReq) (*authentication_pb.VerifyCaptchaResp, error) {
	ok := l.svcCtx.CaptchaStore.Verify(in.CaptchaId, in.Solution, true)
	return &authentication_pb.VerifyCaptchaResp{
		Success: ok,
	}, nil
}
