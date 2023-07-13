package client

import (
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/zrpc"
)

type LuckydrawRpc struct {
	luckydraw_pb.BaseClient
	luckydraw_pb.LuckydrawClient
}

func NewLuckydrawRpc(cli zrpc.Client) *LuckydrawRpc {
	return &LuckydrawRpc{
		BaseClient:      luckydraw_pb.NewBaseClient(cli.Conn()),
		LuckydrawClient: luckydraw_pb.NewLuckydrawClient(cli.Conn()),
	}
}
