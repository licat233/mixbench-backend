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
	conn := cli.Conn()
	return &AdminRpc{
		BaseClient:  admin_pb.NewBaseClient(conn),
		AdminClient: admin_pb.NewAdminClient(conn),
	}
}
