/*
 * @Author: licat
 * @Date: 2022-11-17 14:55:32
 * @LastEditors: licat
 * @LastEditTime: 2022-12-30 11:49:32
 * @Description: licat233@gmail.com
 */
package upload

import (
	"context"
	"fmt"
	"net/http"
	"path"

	"mixbench/common/filex"
	"mixbench/common/respx"
	"mixbench/service/v1/admin/api/internal/svc"

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

func (l *UploadAvatarLogic) UploadAvatar() (resp any, err error) {
	saveDir := path.Join(l.svcCtx.Config.StaticConfig.Root, "avatar")
	err = filex.SaveHttpImg(l.ctx, l.r, saveDir, "avatar")
	if err != nil {
		return nil, err
	}
	url := fmt.Sprint("//", l.svcCtx.Config.StaticConfig.Domain, saveDir)
	return respx.DefaultSingleResp(url, nil)
}
