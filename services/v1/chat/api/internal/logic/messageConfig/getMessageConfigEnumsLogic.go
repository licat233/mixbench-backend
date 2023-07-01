package messageConfig

import (
	"context"

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

func (l *GetMessageConfigEnumsLogic) GetMessageConfigEnums(req *types.GetMessageConfigEnumsReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
