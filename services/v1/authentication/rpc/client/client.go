package client

import (
	"mixbench/services/v1/authentication/rpc/authentication_pb"

	"github.com/zeromicro/go-zero/zrpc"
)

type AuthenticationRpc struct {
	authentication_pb.BaseClient
	authentication_pb.AuthenticationClient
}

func NewAuthenticationRpc(cli zrpc.Client) *AuthenticationRpc {
	return &AuthenticationRpc{
		BaseClient:           authentication_pb.NewBaseClient(cli.Conn()),
		AuthenticationClient: authentication_pb.NewAuthenticationClient(cli.Conn()),
	}
}
