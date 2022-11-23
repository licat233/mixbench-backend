package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Mode  string
	Debug bool

	service.ServiceConf

	Redis redis.RedisConf

	//kq: pub sub
	GoodluckConf kq.KqConf

	//rpc
	LuckdrawRpc zrpc.RpcClientConf

	//mysql
	Mysql struct {
		DataSource string
	}

	//cache
	CacheRedis cache.CacheConf
}
