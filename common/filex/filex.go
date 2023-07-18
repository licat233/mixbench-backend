package filex

import (
	"context"
	"errors"
	"fmt"
	"io"
	"mixbench/common/ctxdata"
	"mixbench/common/errorx"
	"mixbench/common/utils"
	"net/http"
	"os"
	"path"
	"regexp"
	"time"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
)

const MaxFileSize = 10 << 20 // 10 MB

/**
 * @description: 保存http上传的图片，返回图片name
 * @param {*http.Request} r http请求
 * @param {string} saveDir 保存目录名
 * @param {string} prefix 文件名前缀
 * @return {error} 注意:此处的错误为errorx错误
 */
func SaveHttpImg(ctx context.Context, r *http.Request, saveDir, prefix string) (string, error) {
	// 如果 path 路径不存在，会有 err，然后通过 IsNotExist 判定文件路径是否存在，如果 true 则不存在，注意用 os.ModePerm 这样文件是可以写入的
	if _, err := os.Stat(saveDir); os.IsNotExist(err) {
		// mkdir 创建目录，mkdirAll 可创建多层级目录
		os.MkdirAll(saveDir, os.ModePerm)
	}
	r.ParseMultipartForm(MaxFileSize)
	file, header, err := r.FormFile("file")
	if err != nil {
		return "", errorx.ExternalError(err, "上傳文件失敗")
	}
	defer file.Close()
	_, filetype := utils.FileNameAndType(header.Filename)
	if filetype == "" {
		//不存在文件格式
		return "", errorx.ExternalError(nil, "未知文件類型")
	}
	ok, err := ImgTypeVerify(ctx, filetype)
	if !ok {
		return "", err
	}

	date_at := time.Now().Local().Format("20060102")
	uid := ctxdata.GetUidFromCtx(ctx)
	uuid := uuid.New()
	filename := fmt.Sprintf("%s_%s_%d_%d.%s", prefix, date_at, uid, uuid.ID(), filetype)
	filepath := path.Join(saveDir, filename)
	if utils.FileExist(filepath) { //文件已经存在
		return "", errorx.ExternalError(nil, errors.New("上傳異常，請重试"))
	}

	tempFile, err := os.Create(filepath)
	if err != nil {
		msg := fmt.Errorf("文件保存失敗:%v", err)
		logx.WithContext(ctx).Error(msg)
		return "", errorx.InternalError(msg)
	}
	defer tempFile.Close()
	_, err = io.Copy(tempFile, file)
	if err != nil {
		msg := fmt.Errorf("文件复制失敗:%v", err)
		logx.WithContext(ctx).Error(msg)
		return "", errorx.InternalError(msg)
	}
	return filename, nil
}

// ImgTypeVerify 图片类型检查，返回errorx错误
func ImgTypeVerify(ctx context.Context, fileType string) (bool, error) {
	ok, err := regexp.MatchString("^(jpg|png|jpeg|gif|ico|svg|webp|bmp)?$", fileType)
	if err != nil {
		logx.WithContext(ctx).Errorf("regexp.MatchString出错%+v", err)
		return false, errorx.InternalError(fmt.Errorf("regexp.MatchString出错%+v", err), nil)
	}
	if !ok {
		return false, errorx.ExternalError(nil, errors.New("不支持當前圖片格式，受支持的格式：jpg|png|jpeg|gif|ico|svg|webp|bmp"))
	}
	return true, nil
}
