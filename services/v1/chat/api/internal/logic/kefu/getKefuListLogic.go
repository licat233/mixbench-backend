package kefu

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

type GetKefuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetKefuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetKefuListLogic {
	return &GetKefuListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetKefuListLogic) GetKefuList(req *types.GetKefuListReq) (any, error) {
	in := &model.Kefu{
		Id:        req.Id,
		Uuid:      req.Uuid,
		ProjectId: req.ProjectId,
		Passport:  req.Passport,
		Username:  req.Username,
		Password:  req.Password,
		Status:    req.Status,
	}
	mdList, total, err := l.svcCtx.KefuModel.FindList(l.ctx, req.PageSize, req.Page, req.Keyword, in)
	if err != nil {
		l.Logger.Error("failed to query kefu list, error: ", err)
		return nil, errorx.InternalError(err)
	}
	data := dataconv.MdKefus2ApiKefus(mdList)

	return respx.DefaultListResp(data, total, req.PageSize, req.Page, nil)
}
