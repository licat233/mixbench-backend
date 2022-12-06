package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Mode       string
	Debug      bool
	CaptchaRpc zrpc.RpcClientConf
	AdminRpc   zrpc.RpcClientConf
	Redis      struct {
		Addr string
		Type string
		Pass string
		DB   int
	}
	RootDir string
}
