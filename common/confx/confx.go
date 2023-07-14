package confx

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

//通用的一些配置结构

type RpcClientConf struct {
	AuthenticationRpc zrpc.RpcClientConf
	AdminRpc          zrpc.RpcClientConf
	LuckydrawRpc      zrpc.RpcClientConf
}

type ApiConf struct {
	StaticConfig struct {
		Domain string
		Root   string
	}
	RootDir string
}

type RpcConf struct {
	Mysql struct {
		DataSource string
	}
	CacheRedis cache.CacheConf
}

type MqConf struct {
	//kq: pub sub
	KqConf kq.KqConf

	Redis redis.RedisConf
}
