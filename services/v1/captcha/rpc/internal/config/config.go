package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	Debug bool
	Email struct {
		Username string
		Secret   string
		Host     string
		Port     int
		Subject  string
		Receiver []string
	}
}
