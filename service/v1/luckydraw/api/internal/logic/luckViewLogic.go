package logic

import (
	"context"
	"mixbench/service/v1/luckydraw/api/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LuckViewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLuckViewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LuckViewLogic {
	return &LuckViewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LuckViewLogic) LuckView() error {
	// todo: add your logic here and delete this line

	return nil
}
