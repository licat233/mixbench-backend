/*
 * @Author: licat
 * @Date: 2022-11-17 14:55:32
 * @LastEditors: licat
 * @LastEditTime: 2022-12-30 16:46:16
 * @Description: licat233@gmail.com
 */
package upload

import (
	"context"
	"net/http"

	"mixbench/service/v1/admin/api/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadLogoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewUploadLogoLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *UploadLogoLogic {
	return &UploadLogoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *UploadLogoLogic) UploadLogo() (any, error) {
	return UploadImage(l.ctx, l.r, l.svcCtx.Config, "logo")
}
