package ipblack

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/common/respx"
	"mixbench/services/v1/chat/api/internal/svc"
	"mixbench/services/v1/chat/api/internal/types"
	"mixbench/services/v1/chat/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddIpblackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddIpblackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddIpblackLogic {
	return &AddIpblackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddIpblackLogic) AddIpblack(req *types.AddIpblackReq) (any, error) {
	in := &model.Ipblack{
		Ip:     req.Ip,
		KefuId: req.KefuId,
	}
	if _, err := l.svcCtx.IpblackModel.Insert(l.ctx, in); err != nil {
		l.Logger.Error("failed to insert ipblack, error: ", err)
		return nil, errorx.InternalError(err)
	}

	return respx.DefaultStatusResp(nil)
}
