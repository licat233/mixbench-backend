package main

import (
	"flag"
	"fmt"
	"log"
	"mixbench/services/v1/admin/api/internal/config"
	"mixbench/services/v1/admin/api/internal/handler"
	"mixbench/services/v1/admin/api/internal/svc"
	"net/http"

	"mixbench/common/errorx"
	"mixbench/common/utils"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/admin.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	//加载配置文件
	conf.MustLoad(*configFile, &c, conf.UseEnv())

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	staticDir := utils.GenStaticServiceDir("admin")
	if err := utils.CreateDir(staticDir); err != nil {
		log.Fatalln(err)
	}
	urlPrefix := utils.GenStaticServiceUrl("admin") //"/v1/api/admin/static/:file"
	server.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    urlPrefix + ":file",
		Handler: http.StripPrefix(urlPrefix, http.FileServer(http.Dir(staticDir))).ServeHTTP,
	})

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	//错误处理
	httpx.SetErrorHandlerCtx(errorx.ResponseErrorHandlerCtx)

	fmt.Printf("Starting static server at %s\n", urlPrefix)
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
