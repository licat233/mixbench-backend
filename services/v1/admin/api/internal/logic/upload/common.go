package upload

import (
	"context"
	"fmt"
	"mixbench/common/filex"
	"mixbench/common/respx"
	"mixbench/services/v1/admin/api/internal/config"
	"net/http"
	"path"
)

func UploadImage(ctx context.Context, r *http.Request, config config.Config, category string) (any, error) {
	saveDir := path.Join(config.StaticConfig.Root, category)
	name, err := filex.SaveHttpImg(ctx, r, saveDir, category)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprint("//", path.Join(config.StaticConfig.Domain, category, name))
	return respx.DefaultSingleResp(url, nil)
}
