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

type GetKefuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetKefuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetKefuLogic {
	return &GetKefuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetKefuLogic) GetKefu(req *types.GetKefuReq) (any, error) {
	mdKefu, err := l.svcCtx.KefuModel.FindById(l.ctx, req.Id)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error("failed to find kefu, error: ", err)
		return nil, errorx.InternalError(err)
	}
	data := dataconv.MdKefuToApiKefu(mdKefu)

	return respx.DefaultSingleResp(data, nil)
}
