package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"mixbench/common/errorx"
	"mixbench/common/utils"
	"mixbench/services/v1/authentication/api/internal/config"
	"mixbench/services/v1/authentication/api/internal/handler"
	"mixbench/services/v1/authentication/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/authentication.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	staticDir := utils.GenStaticServiceDir("authentication")
	if err := utils.CreateDir(staticDir); err != nil {
		log.Fatalln(err)
	}
	urlPrefix := utils.GenStaticServiceUrl("authentication")
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
