package ipblack

import (
	"context"

	"mixbench/services/v1/chat/api/internal/svc"
	"mixbench/services/v1/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetIpblackEnumsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetIpblackEnumsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetIpblackEnumsLogic {
	return &GetIpblackEnumsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetIpblackEnumsLogic) GetIpblackEnums(req *types.GetIpblackEnumsReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
