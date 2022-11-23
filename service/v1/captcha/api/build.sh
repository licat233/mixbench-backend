#!/bin/bash
#编译captcha-api服务
set -m
current_path=$(
    cd $(dirname $0)
    pwd
)
cd $current_path
pid=$(ps -ef | grep captcha-api | grep -v grep | awk '{print $2}')
if [ ! -z $pid ]; then
    kill -9 $pid
fi
rm -f captcha-api
go build -o captcha-api captcha.go
echo "captcha-api build finished"
