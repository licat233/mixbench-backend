/*
 * @Author: licat
 * @Date: 2023-02-20 14:50:45
 * @LastEditors: licat
 * @LastEditTime: 2023-02-20 16:05:37
 * @Description: licat233@gmail.com
 */

package luckydraw

import (
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/zrpc"
)

type LuckydrawRpc struct {
	luckydraw_pb.BaseClient
	luckydraw_pb.LuckydrawClient
}

func NewLuckydrawRpc(cli zrpc.Client) LuckydrawRpc {
	return LuckydrawRpc{
		BaseClient:      luckydraw_pb.NewBaseClient(cli.Conn()),
		LuckydrawClient: luckydraw_pb.NewLuckydrawClient(cli.Conn()),
	}
}
