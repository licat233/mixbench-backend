package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Email struct {
		Username string
		Secret   string
		Host     string
		Port     int
		Subject  string
		Receiver []string
	}
	Captcha struct {
		Expire int64 //单位second
	}
	KqConf kq.KqConf
}
