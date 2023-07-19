/*
 * @Author: licat
 * @Date: 2022-11-17 14:55:32
 * @LastEditors: licat
 * @LastEditTime: 2023-02-17 15:43:03
 * @Description: licat233@gmail.com
 */
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

	CacheRedis cache.CacheConf

	AuthenticationRpc zrpc.RpcClientConf
	AdminRpc          zrpc.RpcClientConf
	StaticConfig      struct {
		Domain string
		Root   string
	}
}
