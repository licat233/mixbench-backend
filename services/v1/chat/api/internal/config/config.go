package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Mysql struct {
		DataSource string
	}
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	VisitorAuthSecret string
	CacheRedis        cache.CacheConf
	StaticConfig      struct {
		Domain string
		Root   string
	}

	AuthenticationRpc zrpc.RpcClientConf

	RootDir string
}
