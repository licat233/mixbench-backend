/*
 * @Author: licat
 * @Date: 2023-02-01 16:08:53
 * @LastEditors: licat
 * @LastEditTime: 2023-02-01 16:09:34
 * @Description: licat233@gmail.com
 */
package main

import (
	"flag"
	"fmt"

	"mixbench/common/errorx"
	"mixbench/service/v1/luckydraw/rpc/internal/config"
	"mixbench/service/v1/luckydraw/rpc/internal/server"
	"mixbench/service/v1/luckydraw/rpc/internal/svc"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

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
	//配置debug模式
	errorx.Debug = c.Debug

	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		luckydraw_pb.RegisterLuckydrawServer(grpcServer, server.NewLuckydrawServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
