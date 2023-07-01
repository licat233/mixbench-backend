package kefu

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/common/respx"
	"mixbench/services/v1/chat/api/internal/svc"
	"mixbench/services/v1/chat/api/internal/types"
	"mixbench/services/v1/chat/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutKefuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPutKefuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutKefuLogic {
	return &PutKefuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutKefuLogic) PutKefu(req *types.PutKefuReq) (any, error) {
	in := &model.Kefu{
		Id:       req.Id,
		Uuid:     req.Uuid,
		Passport: req.Passport,
		Username: req.Username,
		Password: req.Password,
		Status:   req.Status,
	}
	if err := l.svcCtx.KefuModel.Update(l.ctx, in); err != nil {
		l.Logger.Error("failed to update kefu, error: ", err)
		return nil, errorx.InternalError(err)
	}

	return respx.DefaultStatusResp(nil)
}
