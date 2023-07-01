package logic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/common/respx"
	"mixbench/services/v1/chat/api/internal/svc"
	"mixbench/services/v1/chat/api/internal/types"
	"mixbench/services/v1/chat/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveMsgConfLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveMsgConfLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveMsgConfLogic {
	return &SaveMsgConfLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveMsgConfLogic) SaveMsgConf(req *types.SaveMsgConfReq) (any, error) {
	if len(req.ConfData) == 0 {
		return respx.DefaultStatusResp(nil)
	}
	for _, item := range req.ConfData {
		err := l.svcCtx.MessageConfigModel.UpdateContent(l.ctx, item.Id, item.ConfValue)
		if err != nil && err != model.ErrNotFound {
			l.Logger.Error(err)
			return nil, errorx.InternalError(err)
		}
	}
	return respx.DefaultStatusResp(nil)
}
