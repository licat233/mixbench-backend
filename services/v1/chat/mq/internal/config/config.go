package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	Redis redis.RedisConf

	//kq: pub sub
	KqConf kq.KqConf

	//mysql
	Mysql struct {
		DataSource string
	}

	//cache
	CacheRedis cache.CacheConf
}
