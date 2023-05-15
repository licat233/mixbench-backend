#!/bin/bash

set -m
current_path=$(
    cd $(dirname $0)
    pwd
)
cd $current_path

pid=$(ps -ef | grep captcha-rpc | grep -v grep | awk '{print $2}')
if [ ! -z $pid ]; then
    kill -9 $pid
fi
#编译captcha-rpc服务
rm -f captcha-rpc
go build -o captcha-rpc captcha.go
echo "captcha-rpc build finished"
