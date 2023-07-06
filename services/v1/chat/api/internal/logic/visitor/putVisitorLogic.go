package visitor

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/common/respx"
	"mixbench/services/v1/chat/api/internal/svc"
	"mixbench/services/v1/chat/api/internal/types"
	"mixbench/services/v1/chat/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutVisitorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPutVisitorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutVisitorLogic {
	return &PutVisitorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutVisitorLogic) PutVisitor(req *types.PutVisitorReq) (any, error) {
	in := &model.Visitor{
		Id:       req.Id,
		Uuid:     req.Uuid,
		Name:     req.Name,
		Ip:       req.Ip,
		KefuId:   req.KefuId,
		Status:   req.Status,
		Referer:  req.Referer,
		Region:   req.Region,
		City:     req.City,
		Platform: req.Platform,
		Os:       req.Os,
		Browser:  req.Browser,
		Language: req.Language,
	}
	if err := l.svcCtx.VisitorModel.Update(l.ctx, in); err != nil {
		l.Logger.Error("failed to update visitor, error: ", err)
		return nil, errorx.InternalError(err)
	}

	return respx.DefaultStatusResp(nil)
}
