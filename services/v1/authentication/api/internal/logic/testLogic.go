package logic

import (
	"context"
	"errors"

	"mixbench/common/errorx"
	"mixbench/common/respx"
	"mixbench/services/v1/authentication/api/internal/svc"
	"mixbench/services/v1/authentication/api/internal/types"

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

func (l *TestLogic) Test(req *types.TestReq) (any, error) {
	res, err := l.svcCtx.RedisClient.GetCtx(l.ctx, "licat233")
	if err != nil {
		return nil, errorx.InternalError(err)
	}
	if res == "" {
		return nil, errors.New("不存在!!")
	}
	return respx.NewData(res), nil
}
