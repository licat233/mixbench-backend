package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Redis struct {
		Addr string
		Type string
		Pass string
		DB   int
	}
	Mysql struct {
		DataSource string
	}
	CacheRedis cache.CacheConf

	RootDir      string
	StaticConfig struct {
		Domain string
		Root   string
	}
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	AuthenticationRpc zrpc.RpcClientConf
	AdminRpc          zrpc.RpcClientConf
}
