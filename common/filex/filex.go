package filex

import (
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

	"github.com/zeromicro/go-zero/core/logx"
)

const MaxFileSize = 10 << 20 // 10 MB

/** SaveHttpImg 保存http上传图片，返回图片url
 * r *http.Request http请求
 * saveDir string 保存目录
 * prefix string 文件名前缀
 */
func SaveHttpImg(r *http.Request, saveDir, prefix string) error {
	r.ParseMultipartForm(MaxFileSize)
	file, header, err := r.FormFile("file")
	if err != nil {
		return errorx.ExternalError(err, errors.New("上傳文件失敗"))
	}
	defer file.Close()
	_, filetype, err := FileNameAndType(header.Filename)
	if err != nil {
		return err
	}
	ok, err := ImgTypeVerify(filetype)
	if !ok {
		return err
	}
	ctx := r.Context()
	uid := ctxdata.GetUidFromCtx(ctx)
	filename := GenerateImgName(fmt.Sprintf("%s_%d_", prefix, uid))
	name := fmt.Sprint(filename, ".", filetype)
	filepath := path.Join(saveDir, name)
	if FileExist(filepath) {
		return errorx.ExternalError(nil, errors.New("上傳異常，請重试"))
	}
	tempFile, err := os.Create(filepath)
	if err != nil {
		logx.Errorf("文件保存失敗:%v", err)
		return errorx.InternalError(err)
	}
	defer tempFile.Close()
	_, err = io.Copy(tempFile, file)
	return err
}

// ImgTypeVerify 图片类型检查，返回errorx错误
func ImgTypeVerify(fileType string) (bool, error) {
	ok, err := regexp.MatchString("^(jpg|png|jpeg|gif|ico|svg|webp|bmp)?$", fileType)
	if err != nil {
		logx.Errorf("regexp.MatchString出错%+v", err)
		return false, errorx.InternalError(err, nil)
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
	return fmt.Sprintf("%s%d", prefix, time.Now().UnixNano())
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
