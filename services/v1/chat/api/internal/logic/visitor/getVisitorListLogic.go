package visitor

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

type GetVisitorListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetVisitorListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVisitorListLogic {
	return &GetVisitorListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetVisitorListLogic) GetVisitorList(req *types.GetVisitorListReq) (any, error) {
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
	mdList, total, err := l.svcCtx.VisitorModel.FindList(l.ctx, req.PageSize, req.Page, req.Keyword, in)
	if err != nil {
		l.Logger.Error("failed to query visitor list, error: ", err)
		return nil, errorx.InternalError(err)
	}
	data := dataconv.MdVisitors2ApiVisitors(mdList)

	return respx.DefaultListResp(data, total, req.PageSize, req.Page, nil)
}
