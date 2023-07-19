package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Mysql struct {
		DataSource string
	}
	VisitorAuthSecret string
	CacheRedis        cache.CacheConf

	StaticConfig struct {
		Domain string
		Root   string
	}

	AuthenticationRpc zrpc.RpcClientConf
	AdminRpc          zrpc.RpcClientConf

	RedisDB int //该redis库用于消息分发
	Redis   *redis.RedisKeyConf
}
