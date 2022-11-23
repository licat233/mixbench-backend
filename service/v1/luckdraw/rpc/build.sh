#!/bin/bash
set -m
current_path=$(
    cd $(dirname $0)
    pwd
)
cd $current_path
pid=$(ps -ef | grep luckdraw-rpc | grep -v grep | awk '{print $2}')
if [ ! -z $pid ]; then
    kill -9 $pid
fi
rm -f luckdraw-rpc
go build -o luckdraw-rpc luckdraw.go