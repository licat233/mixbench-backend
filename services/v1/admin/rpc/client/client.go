package client

import (
	"mixbench/services/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/zrpc"
)

type AdminRpc struct {
	admin_pb.BaseClient
	admin_pb.AdminClient
}

func NewAdminRpc(cli zrpc.Client) *AdminRpc {
	return &AdminRpc{
		BaseClient:  admin_pb.NewBaseClient(cli.Conn()),
		AdminClient: admin_pb.NewAdminClient(cli.Conn()),
	}
}
