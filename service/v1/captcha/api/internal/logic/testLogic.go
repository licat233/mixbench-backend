package logic

import (
	"context"

	"mixbench/common/respx"
	"mixbench/service/v1/captcha/api/internal/svc"
	"mixbench/service/v1/captcha/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestLogic {
	return &TestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TestLogic) Test(req *types.TestReq) (resp any, err error) {
	data := "captcha service"
	return respx.DefaultSingleResp(data, nil)
}
