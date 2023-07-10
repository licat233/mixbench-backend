package messageConfig

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/common/respx"
	"mixbench/services/v1/chat/api/internal/logic/dataconv"
	"mixbench/services/v1/chat/api/internal/svc"
	"mixbench/services/v1/chat/api/internal/types"
	"mixbench/services/v1/chat/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMessageConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMessageConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMessageConfigLogic {
	return &GetMessageConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMessageConfigLogic) GetMessageConfig(req *types.GetMessageConfigReq) (any, error) {
	mdMessageConfig, err := l.svcCtx.MessageConfigModel.FindById(l.ctx, req.Id)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error("failed to find messageConfig, error: ", err)
		return nil, errorx.InternalError(err)
	}
	data := dataconv.MdMessageConfigToApiMessageConfig(mdMessageConfig)

	return respx.DefaultSingleResp(data, nil)
}
