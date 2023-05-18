package upload

import (
	"context"
	"net/http"

	"mixbench/services/v1/admin/api/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadAvatarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewUploadAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *UploadAvatarLogic {
	return &UploadAvatarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *UploadAvatarLogic) UploadAvatar() (any, error) {
	return UploadImage(l.ctx, l.r, l.svcCtx.Config, "avatar")
}
