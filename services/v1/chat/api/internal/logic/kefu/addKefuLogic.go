package kefu

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/common/respx"
	"mixbench/common/uniqueid"
	"mixbench/services/v1/chat/api/internal/svc"
	"mixbench/services/v1/chat/api/internal/types"
	"mixbench/services/v1/chat/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddKefuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddKefuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddKefuLogic {
	return &AddKefuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddKefuLogic) AddKefu(req *types.AddKefuReq) (any, error) {
	in := &model.Kefu{
		Uuid:     uniqueid.NewUUID(), // 这里的uniqueid包，请自己定义
		Passport: req.Passport,
		Username: req.Username,
		Password: req.Password,
		Status:   req.Status,
	}
	if _, err := l.svcCtx.KefuModel.Insert(l.ctx, in); err != nil {
		l.Logger.Error("failed to insert kefu, error: ", err)
		return nil, errorx.InternalError(err)
	}

	return respx.DefaultStatusResp(nil)
}
