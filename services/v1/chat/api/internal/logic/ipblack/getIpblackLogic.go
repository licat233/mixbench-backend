package ipblack

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

type GetIpblackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetIpblackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetIpblackLogic {
	return &GetIpblackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetIpblackLogic) GetIpblack(req *types.GetIpblackReq) (any, error) {
	mdIpblack, err := l.svcCtx.IpblackModel.FindOne(l.ctx, req.Id)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error("failed to find ipblack, error: ", err)
		return nil, errorx.InternalError(err)
	}
	data := dataconv.MdIpblackToApiIpblack(mdIpblack)

	return respx.DefaultSingleResp(data, nil)
}
