package messageConfig

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

type GetMessageConfigListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMessageConfigListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMessageConfigListLogic {
	return &GetMessageConfigListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMessageConfigListLogic) GetMessageConfigList(req *types.GetMessageConfigListReq) (any, error) {
	in := &model.MessageConfig{
		Id:      req.Id,
		Name:    req.Name,
		Key:     req.Key,
		Content: req.Content,
	}
	mdList, total, err := l.svcCtx.MessageConfigModel.FindList(l.ctx, req.PageSize, req.Page, req.Keyword, in)
	if err != nil {
		l.Logger.Error("failed to query messageConfig list, error: ", err)
		return nil, errorx.InternalError(err)
	}
	data := dataconv.MdMessageConfigs2ApiMessageConfigs(mdList)

	return respx.DefaultListResp(data, total, req.PageSize, req.Page, nil)
}
