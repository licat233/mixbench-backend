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

type GetVisitorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetVisitorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVisitorLogic {
	return &GetVisitorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetVisitorLogic) GetVisitor(req *types.GetVisitorReq) (any, error) {
	mdVisitor, err := l.svcCtx.VisitorModel.FindById(l.ctx, req.Id)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error("failed to find visitor, error: ", err)
		return nil, errorx.InternalError(err)
	}
	data := dataconv.MdVisitorToApiVisitor(mdVisitor)

	return respx.DefaultSingleResp(data, nil)
}
