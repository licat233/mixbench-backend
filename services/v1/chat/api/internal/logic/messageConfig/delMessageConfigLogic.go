package messageConfig

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/common/respx"
	"mixbench/services/v1/chat/api/internal/svc"
	"mixbench/services/v1/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelMessageConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelMessageConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelMessageConfigLogic {
	return &DelMessageConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelMessageConfigLogic) DelMessageConfig(req *types.DelMessageConfigReq) (any, error) {
	if err := l.svcCtx.MessageConfigModel.Delete(l.ctx, req.Id); err != nil {
		l.Logger.Error("failed to delete messageConfig, error: ", err)
		return nil, errorx.InternalError(err)
	}

	return respx.DefaultStatusResp(nil)
}
