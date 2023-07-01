package main

import (
	"flag"
	"fmt"
	"net/http"
	"path"

	"mixbench/common/errorx"
	"mixbench/common/filex"
	"mixbench/services/v1/chat/api/internal/config"
	"mixbench/services/v1/chat/api/internal/handler"
	"mixbench/services/v1/chat/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/chat.yaml", "the config file")
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

	//配置静态资源服务
	staticPath := path.Join(c.RootDir, "static")
	if err := filex.CreateDir(staticPath); err != nil {
		panic(err)
	}
	server.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    "/v1/static/chat/:file",
		Handler: http.StripPrefix("/v1/static/chat/", http.FileServer(http.Dir(staticPath))).ServeHTTP,
	})

	ctx := svc.NewServiceContext(c)

	defer ctx.Hub.Stop()
	go ctx.Hub.Run()

	handler.RegisterHandlers(server, ctx)

	//错误处理
	httpx.SetErrorHandlerCtx(errorx.ResponseErrorHandlerCtx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
