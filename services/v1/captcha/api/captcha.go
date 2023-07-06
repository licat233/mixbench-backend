package main

import (
	"flag"
	"fmt"
	"mixbench/services/v1/captcha/api/internal/config"
	"mixbench/services/v1/captcha/api/internal/handler"
	"mixbench/services/v1/captcha/api/internal/svc"
	"net/http"
	"path"

	"mixbench/common/errorx"
	"mixbench/common/filex"
	"mixbench/common/utils"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/captcha.yaml", "the config file")
var rootDir = flag.String("d", "./", "the project root path")

func main() {
	flag.Parse()

	var c config.Config
	//加载配置文件
	conf.MustLoad(*configFile, &c, conf.UseEnv())
	c.RootDir = *rootDir
	//配置debug模式
	errorx.Debug = c.Debug

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	staticPath := path.Join(c.RootDir, "static")
	if err := filex.CreateDir(staticPath); err != nil {
		panic(err)
	}
	urlPrefix := utils.GenStaticServiceUrl("captcha")
	server.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    urlPrefix + ":file",
		Handler: http.StripPrefix(urlPrefix, http.FileServer(http.Dir(staticPath))).ServeHTTP,
	})

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	//错误处理
	httpx.SetErrorHandlerCtx(errorx.ResponseErrorHandlerCtx)

	fmt.Printf("Starting static server at %s\n", urlPrefix)
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
