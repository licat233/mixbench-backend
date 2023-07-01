package chatRecord

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

type GetChatRecordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetChatRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetChatRecordLogic {
	return &GetChatRecordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetChatRecordLogic) GetChatRecord(req *types.GetChatRecordReq) (any, error) {
	mdChatRecord, err := l.svcCtx.ChatRecordModel.FindOne(l.ctx, req.Id)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error("failed to find chatRecord, error: ", err)
		return nil, errorx.InternalError(err)
	}
	data := dataconv.MdChatRecordToApiChatRecord(mdChatRecord)

	return respx.DefaultSingleResp(data, nil)
}
