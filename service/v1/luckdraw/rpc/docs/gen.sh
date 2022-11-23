nohup=`goctl rpc protoc luckdraw.proto --go_out=../ --go-grpc_out=../ --zrpc_out=../ -style goZero`
sed -i "" 's/,omitempty//g' "../luckdraw_pb/luckdraw.pb.go"