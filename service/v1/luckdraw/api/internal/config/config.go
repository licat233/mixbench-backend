package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Mode    string
	Debug   bool
	RootDir string
	JwtAuth struct {
		AccessSecret string
	}
	CaptchaRpc  zrpc.RpcClientConf
	LuckdrawRpc zrpc.RpcClientConf
}
