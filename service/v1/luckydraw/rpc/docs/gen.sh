#!/bin/bash
###
# @Author: licat
# @Date: 2023-01-15 14:05:33
 # @LastEditors: licat
 # @LastEditTime: 2023-02-01 16:44:46
# @Description: licat233@gmail.com
###

#进入monitor mode
set -m

current_path=$(
    cd $(dirname $0)
    pwd
)

project_root="${current_path%%mixbench-backend*}mixbench-backend"

if [ ! -d backup ]; then
    mkdir backup
fi

name="luckydraw"
multiple=false

# 如果存在，做个备份
if [ -f ${name}.proto ]; then
    if [ ! -f backup/${name}1.proto ]; then
        cp ${name}.proto backup/${name}1.proto
    elif [ ! -f ./backup/${name}2.proto ]; then
        cp ${name}.proto backup/${name}2.proto
    elif [ ! -f ./backup/${name}3.proto ]; then
        cp ${name}.proto backup/${name}3.proto
    elif [ ! -f ./backup/${name}4.proto ]; then
        cp ${name}.proto backup/${name}4.proto
    elif [ ! -f ./backup/${name}5.proto ]; then
        cp ${name}.proto backup/${name}5.proto
    elif [ -f ./backup/${name}5.proto ]; then
        rm -f backup/${name}1.proto
        mv backup/${name}2.proto backup/${name}1.proto
        mv backup/${name}3.proto backup/${name}2.proto
        mv backup/${name}4.proto backup/${name}3.proto
        mv backup/${name}5.proto backup/${name}4.proto
        cp ${name}.proto backup/${name}5.proto
    else
        cp ${name}.proto backup/${name}1.proto
    fi
fi

${project_root}/deploy/sql2pb/build/sql2pb -schema=mixbench_${name} -package=${name}_proto -go_package=./${name}_pb -service_name=${name} -multiple=${multiple} -filename=${name}.proto

# exit

goctl rpc protoc ${name}.proto --go_out=../ --go-grpc_out=../ --zrpc_out=../ --style=goZero --multiple=${multiple}

if [ $? -ne 0 ]; then
    exit
fi

# macOS环境中需要传入一个文件做备份
sed -i "" 's/,omitempty//g' "../${name}_pb/${name}.pb.go" # for macOS
if [ $? -ne 0 ]; then
    sed -i 's/,omitempty//g' "../${name}_pb/${name}.pb.go" # for linux
fi
