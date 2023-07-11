package main

import (
	"flag"
	"fmt"

	"mixbench/services/v1/luckydraw/rpc/internal/config"
	baseluckydrawServer "mixbench/services/v1/luckydraw/rpc/internal/server/base"
	luckydrawServer "mixbench/services/v1/luckydraw/rpc/internal/server/luckydraw"
	"mixbench/services/v1/luckydraw/rpc/internal/svc"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/luckydraw.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	//加载配置文件
	conf.MustLoad(*configFile, &c, conf.UseEnv())

	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		luckydraw_pb.RegisterBaseServer(grpcServer, baseluckydrawServer.NewBaseServer(ctx))
		luckydraw_pb.RegisterLuckydrawServer(grpcServer, luckydrawServer.NewLuckydrawServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
