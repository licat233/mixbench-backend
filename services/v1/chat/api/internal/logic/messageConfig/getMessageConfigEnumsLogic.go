package messageConfig

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/common/respx"
	"mixbench/services/v1/chat/api/internal/svc"
	"mixbench/services/v1/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMessageConfigEnumsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMessageConfigEnumsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMessageConfigEnumsLogic {
	return &GetMessageConfigEnumsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMessageConfigEnumsLogic) GetMessageConfigEnums(req *types.GetMessageConfigEnumsReq) (any, error) {
	mdMessageConfigList, err := l.svcCtx.MessageConfigModel.FindAll(l.ctx)
	if err != nil {
		l.Logger.Error("failed to query messageConfig list, error: ", err)
		return nil, errorx.InternalError(err)
	}
	data := make([]*types.Enum, 0)
	for _, v := range mdMessageConfigList {
		data = append(data, &types.Enum{
			Label: v.Name,
			Value: v.Id,
		})
	}

	return respx.DefaultSingleResp(data, nil)
}
