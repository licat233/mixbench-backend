/*
 * @Author: licat
 * @Date: 2022-12-30 16:40:59
 * @LastEditors: licat
 * @LastEditTime: 2022-12-30 16:44:59
 * @Description: licat233@gmail.com
 */
package upload

import (
	"context"
	"fmt"
	"mixbench/common/filex"
	"mixbench/common/respx"
	"mixbench/service/v1/admin/api/internal/config"
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
