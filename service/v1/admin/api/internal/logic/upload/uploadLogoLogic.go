package upload

import (
	"context"
	"fmt"
	"net/http"
	"path"

	"mixbench/common/errorx"
	"mixbench/common/filex"
	"mixbench/common/respx"
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

func (l *UploadLogoLogic) UploadLogo() (resp any, err error) {
	saveDir := path.Join(l.svcCtx.Config.StaticConfig.Root, "logo")
	err = filex.SaveHttpImg(l.r, saveDir, "logo")
	if err != nil {
		logx.Error(err)
		return nil, errorx.InternalError(err)
	}
	url := fmt.Sprint("//", l.svcCtx.Config.StaticConfig.Domain, saveDir)
	return respx.DefaultSingleResp(url, nil)
}
