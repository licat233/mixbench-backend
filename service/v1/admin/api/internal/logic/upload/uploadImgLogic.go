/*
 * @Author: licat
 * @Date: 2022-11-17 14:55:32
 * @LastEditors: licat
 * @LastEditTime: 2022-12-30 16:46:09
 * @Description: licat233@gmail.com
 */
package upload

import (
	"context"
	"net/http"

	"mixbench/service/v1/admin/api/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadImgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewUploadImgLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *UploadImgLogic {
	return &UploadImgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *UploadImgLogic) UploadImg() (any, error) {
	return UploadImage(l.ctx, l.r, l.svcCtx.Config, "img")
}
