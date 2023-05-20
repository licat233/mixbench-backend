/*
 * @Author: licat
 * @Date: 2022-11-17 14:55:33
 * @LastEditors: licat
 * @LastEditTime: 2022-12-30 16:04:57
 * @Description: licat233@gmail.com
 */
package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Debug      bool
	CaptchaRpc zrpc.RpcClientConf
	AdminRpc   zrpc.RpcClientConf
	Redis      struct {
		Addr string
		Type string
		Pass string
		DB   int
	}
	RootDir      string
	StaticConfig struct {
		Domain string
		Root   string
	}
}
