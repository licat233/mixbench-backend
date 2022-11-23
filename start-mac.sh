#!/bin/bash
set -m

#判断brew是否存在
if false && command -v brew >/dev/null 2>&1; then
    #启动redis
    brew services start redis

    #启动etcd
    brew services start etcd

    #启动kafka
    brew services start kafka

    #启动nginx
    brew services start nginx
fi

current_path=$(
    cd $(dirname $0)
    pwd
)

cd "$current_path/deploy/script/mac"

chmod a+x ./gen.sh
chmod a+x ./build.sh
chmod a+x ./run.sh
chmod a+x ./start.sh

./start.sh $1 $2 $3
