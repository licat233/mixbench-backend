/*
 * @Author: licat
 * @Date: 2022-11-17 14:55:33
 * @LastEditors: licat
 * @LastEditTime: 2022-12-30 16:05:05
 * @Description: licat233@gmail.com
 */
package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Mode    string
	Debug   bool
	RootDir string
	Auth    struct {
		AccessSecret string
	}
	CaptchaRpc   zrpc.RpcClientConf
	LuckydrawRpc zrpc.RpcClientConf
	StaticConfig struct {
		Domain string
		Root   string
	}
}
