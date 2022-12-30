/*
 * @Author: licat
 * @Date: 2022-11-17 14:55:03
 * @LastEditors: licat
 * @LastEditTime: 2022-12-30 17:01:45
 * @Description: licat233@gmail.com
 */
package filex

import (
	"context"
	"errors"
	"fmt"
	"io"
	"mixbench/common/ctxdata"
	"mixbench/common/errorx"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
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
		return "", errorx.ExternalError(err, errors.New("上傳文件失敗"))
	}
	defer file.Close()
	_, filetype, err := FileNameAndType(header.Filename)
	if err != nil {
		return "", err
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
	if FileExist(filepath) { //文件已经存在
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

// FilenameValid 文件名检查，返回原生error
func FilenameValidate(filename string) error {
	n := len(filename)
	if n == 0 {
		return errors.New("該文件沒有名字")
	}
	if n > 64 {
		return errors.New("文件名過長")
	}
	if ok, _ := regexp.MatchString("^.{1,64}\\.[A-Za-z0-9]{2,5}?$", filename); !ok {
		return errors.New("未知文件類型")
	}
	return nil
}

// FileSize 获取文件大小
func FileSize(filePath string) (int64, bool) {
	fi, err := os.Lstat(filePath)
	if err == nil {
		return fi.Size(), !fi.IsDir()
	}
	return fi.Size(), !os.IsNotExist(err)
}

// FileExist 判断文件是否存在
func FileExist(path string) bool {
	fi, err := os.Lstat(path)
	if err == nil {
		return !fi.IsDir()
	}
	return !os.IsNotExist(err)
}

// DirExist 判断文件夹是否存在
func DirExist(path string) bool {
	fi, err := os.Lstat(path)
	if err == nil {
		return fi.IsDir()
	}
	return !os.IsNotExist(err)
}

// FileNameAndType 获取文件名与类型，同时会检查文件名是否合法，或者是否存在
// 返回errorx错误
func FileNameAndType(filename string) (fileName string, fileType string, err error) {
	err = FilenameValidate(filename)
	if err != nil {
		err = errorx.ExternalError(nil, err)
		return
	}
	index := strings.LastIndex(filename, ".")
	if index == -1 || index == len(filename)-1 {
		err = errorx.ExternalError(nil, errors.New("file type does not exist"))
		return
	}
	fileName = filename[:index]
	fileType = filename[index+1:]
	return
}

// GenerateImgName 生成图片名字
func GenerateImgName(prefix string) string {
	return fmt.Sprintf("%s_%d", prefix, time.Now().UnixNano())
}

// 判断文件夹是否存在
func HasDir(path string) (bool, error) {
	_, _err := os.Stat(path)
	if _err == nil {
		return true, nil
	}
	if os.IsNotExist(_err) {
		return false, nil
	}
	return false, _err
}

// 创建文件夹
func CreateDir(path string) error {
	_exist, _err := HasDir(path)
	if _err != nil {
		return fmt.Errorf("获取文件夹异常 -> %v", _err)
	}
	if _exist {
		return nil
	}
	err := os.Mkdir(path, os.ModePerm)
	if err != nil {
		return fmt.Errorf("创建目录异常 -> %v", _err)
	}
	return nil
}
