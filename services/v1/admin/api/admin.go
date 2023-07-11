package main

import (
	"flag"
	"fmt"
	"mixbench/services/v1/admin/api/internal/config"
	"mixbench/services/v1/admin/api/internal/handler"
	"mixbench/services/v1/admin/api/internal/svc"
	"net/http"
	"path"

	"mixbench/common/errorx"
	"mixbench/common/filex"
	"mixbench/common/utils"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/admin.yaml", "the config file")
var rootDir = flag.String("d", "./", "the project root path")

func main() {
	flag.Parse()

	var c config.Config
	//加载配置文件
	conf.MustLoad(*configFile, &c, conf.UseEnv())
	c.RootDir = *rootDir

	//jwt认证错误处理
	// jwtErrHandler := rest.WithUnauthorizedCallback(func(w http.ResponseWriter, r *http.Request, err error) {
	// 	// 自定义处理返回
	// })

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	staticPath := path.Join(c.RootDir, "static")
	if err := filex.CreateDir(staticPath); err != nil {
		panic(err)
	}
	urlPrefix := utils.GenStaticServiceUrl("admin")
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
