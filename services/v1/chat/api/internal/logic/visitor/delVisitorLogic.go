package visitor

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/common/respx"
	"mixbench/services/v1/chat/api/internal/svc"
	"mixbench/services/v1/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelVisitorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelVisitorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelVisitorLogic {
	return &DelVisitorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelVisitorLogic) DelVisitor(req *types.DelVisitorReq) (any, error) {
	if err := l.svcCtx.VisitorModel.SoftDelete(l.ctx, req.Id); err != nil {
		l.Logger.Error("failed to soft delete visitor, error: ", err)
		return nil, errorx.InternalError(err)
	}

	return respx.DefaultStatusResp(nil)
}
