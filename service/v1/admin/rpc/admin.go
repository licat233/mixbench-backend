package main

import (
	"flag"
	"fmt"

	"mixbench/common/errorx"
	"mixbench/service/v1/admin/rpc/admin_pb"
	"mixbench/service/v1/admin/rpc/internal/config"
	adminServer "mixbench/service/v1/admin/rpc/internal/server/admin"
	baseServer "mixbench/service/v1/admin/rpc/internal/server/base"
	"mixbench/service/v1/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/admin.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	//加载配置文件
	conf.MustLoad(*configFile, &c, conf.UseEnv())
	//配置debug模式
	errorx.Debug = c.Debug

	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		admin_pb.RegisterBaseServer(grpcServer, baseServer.NewBaseServer(ctx))
		admin_pb.RegisterAdminServer(grpcServer, adminServer.NewAdminServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
