#!/bin/bash
set -m
current_path=$(
    cd $(dirname $0)
    pwd
)
cd $current_path
pid=$(ps -ef | grep luckydraw-rpc | grep -v grep | awk '{print $2}')
if [ ! -z $pid ]; then
    kill -9 $pid
fi
rm -f luckydraw-rpc
go build -o luckydraw-rpc luckydraw.go