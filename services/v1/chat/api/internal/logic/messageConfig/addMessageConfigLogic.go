package messageConfig

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/common/respx"
	"mixbench/services/v1/chat/api/internal/svc"
	"mixbench/services/v1/chat/api/internal/types"
	"mixbench/services/v1/chat/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddMessageConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddMessageConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMessageConfigLogic {
	return &AddMessageConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddMessageConfigLogic) AddMessageConfig(req *types.AddMessageConfigReq) (any, error) {
	in := &model.MessageConfig{
		ConfName:  req.ConfName,
		ConfKey:   req.ConfKey,
		ConfValue: req.ConfValue,
	}
	if _, err := l.svcCtx.MessageConfigModel.Insert(l.ctx, in); err != nil {
		l.Logger.Error("failed to insert messageConfig, error: ", err)
		return nil, errorx.InternalError(err)
	}

	return respx.DefaultStatusResp(nil)
}
