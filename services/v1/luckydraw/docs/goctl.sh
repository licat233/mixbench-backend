#!/bin/bash

#进入monitor mode
set -m

current_path=$(
    cd "$(dirname $0)" || exit
    pwd
)

cd "$current_path" || exit

name="luckydraw"

project_root="${current_path%%mixbench-backend*}mixbench-backend"
sql_file="${project_root}/deploy/mysql/${name}.sql"
if ! goctl model mysql ddl --src "$sql_file" -dir ../model --cache --style goZero; then
    exit
fi

# Please install the goctl first
# command: go install github.com/zeromicro/go-zero/tools/goctl@latest
# github: https://github.com/zeromicro/go-zero

if ! goctl api go --api $name.api --dir ../api --style goZero; then
    exit 1
fi

if ! goctl rpc protoc $name.proto --go_out=../rpc --go-grpc_out=../rpc --zrpc_out=../rpc --multiple=true --style goZero; then
    exit 1
fi

filename="../rpc/${name}_pb/${name}.pb.go"

if ! sed -i "" 's/,omitempty//g' $filename; then
    sed -i 's/,omitempty//g' $filename
fi

go mod tidy

goctl api plugin -plugin goctl-swagger="swagger -filename ../api/static/swagger.json" -api $name.api -dir .
goctl api ts -api $name.api -dir ../api/static