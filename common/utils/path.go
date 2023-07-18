package utils

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

var ProjectRoot string = CurrentProjectRootPath()
var ServiceDir string = ServiceFileDir()

// 获取当前项目地址, 使用这个获取项目根目录
func CurrentProjectRootPath() string {
	fullPath := getCurrentAbPath()
	reg := regexp.MustCompile("^.*common")
	matched := reg.FindString(fullPath)
	res := strings.TrimSuffix(matched, "common")
	return res
}

// 获取当前项目地址
func CurrentPath() string {
	return getCurrentAbPath()
}

// 最终方案-全兼容
func getCurrentAbPath() string {
	dir := getCurrentAbPathByExecutable()
	//如果是 go run 启动的
	tmpPath := getTmpDir()
	if tmpPath == "." {
		return getCurrentAbPathByCaller()
	}
	if strings.Contains(dir, getTmpDir()) {
		return getCurrentAbPathByCaller()
	}
	return dir
}

// 获取系统临时目录，兼容go run
func getTmpDir() string {
	dir := os.Getenv("TEMP")
	if dir == "" {
		dir = os.Getenv("TMP")
	}
	res, _ := filepath.EvalSymlinks(dir)
	return res
}

// 获取当前执行文件绝对路径（仅支持go build）
func getCurrentAbPathByExecutable() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}

// 获取当前执行文件绝对路径（仅支持go run）
func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}

// 生成静态服务url规范，使其与nginx配置匹配
func GenStaticServiceUrl(name string) string {
	return fmt.Sprintf("/v1/static/%s/", name)
}

func GenStaticServiceDir(name string) string {
	if os.Getenv("SERVICE_MODE") == "dev" {
		return fmt.Sprintf("services/v1/%s/api/static", name)
	}
	return path.Join(path.Dir(os.Args[0]), "static")
}

// 服务文件所在目录
func ServiceFileDir() string {
	return path.Dir(os.Args[0])
}

// 是否为公开路径
func IsPublicPath(s string) bool {
	// 使用正则表达式匹配 "/v{版本号}/api/{名称}/public" 后面可能有路径的格式
	re := regexp.MustCompile(`/v\d+/api/([^/]+/)?public(/.*)?`)
	ok := re.MatchString(s)
	return ok
}
