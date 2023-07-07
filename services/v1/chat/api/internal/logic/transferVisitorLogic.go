package logic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/common/respx"
	"mixbench/services/v1/chat/api/internal/svc"
	"mixbench/services/v1/chat/api/internal/types"
	"mixbench/services/v1/chat/api/internal/wshub"

	"github.com/zeromicro/go-zero/core/logx"
)

type TransferVisitorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTransferVisitorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TransferVisitorLogic {
	return &TransferVisitorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TransferVisitorLogic) TransferVisitor(req *types.TransferVisitorReq) (any, error) {
	visitorId := req.VisitorId
	if visitorId <= 0 {
		return nil, errorx.ExternalError("visitorId <=0", "visitorId must be greater than 0")
	}
	kefuId := req.KefuId //可以为0或负数，表示无人接待
	// if kefuId <= 0 {
	// 	return nil, errorx.ExternalError("kefuId <=0", "kefuId must be greater than 0")
	// }

	err := wshub.HubSvc.TransferVisitor(visitorId, kefuId)
	if err != nil {
		return nil, err
	}

	return respx.DefaultStatusResp(nil)
}
