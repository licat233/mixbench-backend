/*
 * @Author: licat
 * @Date: 2023-01-30 16:38:35
 * @LastEditors: licat
 * @LastEditTime: 2023-01-30 17:53:06
 * @Description: licat233@gmail.com
 */
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
