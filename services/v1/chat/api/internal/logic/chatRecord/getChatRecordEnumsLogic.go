package chatRecord

import (
	"context"

	"mixbench/services/v1/chat/api/internal/svc"
	"mixbench/services/v1/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetChatRecordEnumsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetChatRecordEnumsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetChatRecordEnumsLogic {
	return &GetChatRecordEnumsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetChatRecordEnumsLogic) GetChatRecordEnums(req *types.GetChatRecordEnumsReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
