package visitor

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/common/respx"
	"mixbench/common/uniqueid"
	"mixbench/services/v1/chat/api/internal/svc"
	"mixbench/services/v1/chat/api/internal/types"
	"mixbench/services/v1/chat/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddVisitorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddVisitorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddVisitorLogic {
	return &AddVisitorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddVisitorLogic) AddVisitor(req *types.AddVisitorReq) (any, error) {
	in := &model.Visitor{
		Uuid:      uniqueid.NewUUID(), // 这里的uniqueid包，请自己定义
		Name:      req.Name,
		Ip:        req.Ip,
		KefuId:    req.KefuId,
		ProjectId: req.ProjectId,
		Status:    req.Status,
		Referer:   req.Referer,
		Region:    req.Region,
		City:      req.City,
		Platform:  req.Platform,
		Os:        req.Os,
		Browser:   req.Browser,
		Language:  req.Language,
	}
	if _, err := l.svcCtx.VisitorModel.Insert(l.ctx, in); err != nil {
		l.Logger.Error("failed to insert visitor, error: ", err)
		return nil, errorx.InternalError(err)
	}

	return respx.DefaultStatusResp(nil)
}
