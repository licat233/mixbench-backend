/*
 * @Author: licat
 * @Date: 2023-02-17 17:24:28
 * @LastEditors: licat
 * @LastEditTime: 2023-02-17 18:12:18
 * @Description: licat233@gmail.com
 */
package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql struct {
		DataSource string
	}
	CacheRedis cache.CacheConf
}
