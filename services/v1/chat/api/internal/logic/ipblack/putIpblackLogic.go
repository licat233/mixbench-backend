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

type PutIpblackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPutIpblackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutIpblackLogic {
	return &PutIpblackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutIpblackLogic) PutIpblack(req *types.PutIpblackReq) (any, error) {
	in := &model.Ipblack{
		Id: req.Id,
		Ip: req.Ip,
	}
	if err := l.svcCtx.IpblackModel.Update(l.ctx, in); err != nil {
		l.Logger.Error("failed to update ipblack, error: ", err)
		return nil, errorx.InternalError(err)
	}

	return respx.DefaultStatusResp(nil)
}
