package logic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/common/respx"
	"mixbench/services/v1/chat/api/internal/svc"
	"mixbench/services/v1/chat/api/internal/types"
	"mixbench/services/v1/chat/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutVisitorCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPutVisitorCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutVisitorCommentLogic {
	return &PutVisitorCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutVisitorCommentLogic) PutVisitorComment(req *types.PutVisitorCommentReq) (any, error) {
	if req.VisitorId < 1 {
		return nil, errorx.ExternalError("visitor id < 1")
	}
	if len(req.Comment) > 100 {
		return nil, errorx.ExternalError("the length of comment is too large")
	}
	mdViditor, err := l.svcCtx.VisitorModel.FindById(l.ctx, req.VisitorId)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.InternalError(err)
	}
	if err == model.ErrNotFound {
		return nil, errorx.New("visitor does not exist")
	}
	mdViditor.Comment = req.Comment
	err = l.svcCtx.VisitorModel.Update(l.ctx, mdViditor)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.InternalError(err)
	}
	return respx.DefaultStatusResp(nil)
}
