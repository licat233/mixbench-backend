goctl rpc protoc admin.proto --go_out=../ --go-grpc_out=../ --zrpc_out=../ -style goZero
sed -i "" 's/,omitempty//g' "../admin_pb/admin.pb.go"