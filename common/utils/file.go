package utils

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

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
func FileNameAndType(filename string) (fileName string, fileType string) {
	filename = strings.TrimSpace(filename)
	parts := strings.Split(filename, ".")
	if len(parts) < 2 {
		fileName = filename
		return
	}
	index := len(parts) - 1
	fileName = strings.Join(parts[:index], ".")
	fileType = parts[index]
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

// 查找文件，以忽略文件名大小写的方式查找匹配
func FindFile(dir string, file string) (string, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return "", err
	}
	for _, fileInfo := range files {
		if strings.EqualFold(fileInfo.Name(), file) {
			return filepath.Join(dir, fileInfo.Name()), nil
		}
	}
	return "", nil
}

func CreateFile(filePath string) error {
	// 获取文件所在的目录
	dir := filepath.Dir(filePath)

	// 创建目录及其父目录
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// 创建文件
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}

	// 关闭文件
	defer func(f *os.File) {
		if f != nil {
			err := f.Close()
			if err != nil {
				return
			}
		}
	}(f)

	return nil
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, fs.ErrNotExist) {
		return false, nil
	}
	return false, err
}

func WriteFile(filename string, data string) error {
	exists, err := PathExists(filename)
	if err != nil {
		return err
	}
	if !exists {
		err := CreateFile(filename)
		if err != nil {
			return err
		}
	}

	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	_, err = f.WriteString(data)
	if err != nil {
		return err
	}
	return nil
}

func ReadFile(filename string) (string, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return TrimSpace(string(b)), nil
}

func TrimSpace(content string) string {
	n := len(content)
	content = strings.TrimSpace(content)
	content = strings.Trim(content, "\n")
	content = strings.Trim(content, "\t")
	for n != len(content) {
		content = strings.TrimSpace(content)
		content = strings.Trim(content, "\n")
		content = strings.Trim(content, "\t")
		n = len(content)
	}
	return content
}

// InsertContentAfter 文件内容追加
// 会清除文件内容前后换行或者空格
func InsertFileContentAfter(filename string, content string) error {
	fileContent, err := ReadFile(filename)
	if err != nil {
		return err
	}
	fileContent = TrimSpace(fileContent)
	fileContent = fmt.Sprintf("%s\n%s", fileContent, content)
	if err := os.WriteFile(filename, []byte(fileContent), 0644); err != nil {
		return err
	}
	return nil
}

// InsertContentAfter2InsertContentAfter2 文件内容追加
// 不会清除文件内容前后换行或者空格
func InsertFileContentAfter2(filename string, content string) error {
	// 打开文件，如果文件不存在则创建
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	content = TrimSpace(content)
	content = fmt.Sprintf("\n%s", content)
	// 写入要追加的内容
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}
