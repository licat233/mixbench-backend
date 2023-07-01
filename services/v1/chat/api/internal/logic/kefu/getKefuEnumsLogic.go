package kefu

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/common/respx"
	"mixbench/services/v1/chat/api/internal/svc"
	"mixbench/services/v1/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetKefuEnumsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetKefuEnumsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetKefuEnumsLogic {
	return &GetKefuEnumsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetKefuEnumsLogic) GetKefuEnums(req *types.GetKefuEnumsReq) (any, error) {
	mdKefuList, err := l.svcCtx.KefuModel.FindAll(l.ctx)
	if err != nil {
		l.Logger.Error("failed to query kefu list, error: ", err)
		return nil, errorx.InternalError(err)
	}
	data := make([]*types.Enum, 0)
	for _, v := range mdKefuList {
		data = append(data, &types.Enum{
			Label: v.Username,
			Value: v.Id,
		})
	}

	return respx.DefaultSingleResp(data, nil)
}
