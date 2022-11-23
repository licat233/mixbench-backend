#!/bin/bash
set -m
current_path=$(
    cd $(dirname $0)
    pwd
)
cd $current_path
pid=$(ps -ef | grep admin-api | grep -v grep | awk '{print $2}')
if [ ! -z $pid ]; then
    kill -9 $pid
fi
rm -f admin-api
go build -o admin-api admin.go
nohup=$(./admin-api &)
echo "执行完毕:${nohup}"
exit
