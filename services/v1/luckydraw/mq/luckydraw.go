package main

import (
	"flag"
	"fmt"
	"log"
	"mixbench/services/v1/luckydraw/mq/internal/config"
	"mixbench/services/v1/luckydraw/mq/internal/listen"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
)

var configFile = flag.String("f", "etc/luckydraw.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	//加载配置文件
	conf.MustLoad(*configFile, &c, conf.UseEnv())

	if err := c.SetUp(); err != nil {
		log.Fatalln(err)
	}

	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()

	for _, mq := range listen.Mqs(c) {
		serviceGroup.Add(mq)
	}
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	serviceGroup.Start()
}
