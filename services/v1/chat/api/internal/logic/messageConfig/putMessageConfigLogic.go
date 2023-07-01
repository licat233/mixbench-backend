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

type PutMessageConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPutMessageConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutMessageConfigLogic {
	return &PutMessageConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutMessageConfigLogic) PutMessageConfig(req *types.PutMessageConfigReq) (any, error) {
	in := &model.MessageConfig{
		Id:      req.Id,
		Name:    req.Name,
		Key:     req.Key,
		Content: req.Content,
	}
	if err := l.svcCtx.MessageConfigModel.Update(l.ctx, in); err != nil {
		l.Logger.Error("failed to update messageConfig, error: ", err)
		return nil, errorx.InternalError(err)
	}

	return respx.DefaultStatusResp(nil)
}
