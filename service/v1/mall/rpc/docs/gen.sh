#!/bin/bash
###
# @Author: licat
# @Date: 2023-01-15 14:05:33
 # @LastEditors: licat
 # @LastEditTime: 2023-01-27 18:20:10
# @Description: licat233@gmail.com
###

#进入monitor mode
set -m

current_path=$(
    cd $(dirname $0)
    pwd
)

project_root="${current_path%%mixbench-backend*}mixbench-backend"

name="mall"

${project_root}/deploy/sql2pb/build/sql2pb -schema=mixbench_${name} -package=${name}_proto -go_package=./${name}_pb -service_name=${name} >${name}.proto

goctl rpc protoc ${name}.proto --go_out=../ --go-grpc_out=../ --zrpc_out=../ --style=goZero --multiple=false

if [ $? -ne 0 ]; then
    exit
fi

# macOS环境中需要传入一个文件做备份
sed -i "" 's/,omitempty//g' "../core_pb/core.pb.go"
if [ $? -ne 0 ]; then
    sed -i 's/,omitempty//g' "../core_pb/core.pb.go"
fi
