package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

// type BrokerItem struct {
// 	Kafka int64
// }

type KqConfig struct {
	Brokers []string
	Topic   string
}

type Config struct {
	zrpc.RpcServerConf
	Mode  string
	Debug bool
	Mysql struct {
		DataSource string
	}
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
	CacheRedis     cache.CacheConf
	KqGoodluckConf KqConfig
}
