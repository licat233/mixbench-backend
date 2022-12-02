#!/bin/bash

#进入monitor mode
set -m

current_path=$(
    cd $(dirname $0)
    pwd
)

cd $current_path
cd ..

domain="mixbench.net"

# 注意：macOS中 bash3.2不支持关联数组，需要升级到bash4+，但不建议这样做，因为Apple仍将版本3作为系统外壳提供
# 故采取此方案

weblist=(
    "mysql,mysql_manage:80"
    "redis,redis_manage:80"
    "jaeger,jaeger:16686"
    "grafana,grafana:3000"
    "kibana,kibana:5601"
    "asynqmon,asynqmon:8080"
    "filebrowser,filebrowser:80"
    "kibana,kafka_manage:9000"
    "elasticsearch,elasticsearch:9200"
)

splitstr() {
    web=$1
    arr=($(echo $web | tr ',' ' '))
    num=${#arr[@]}
    if [ $num -ne 2 ]; then
        echo "error: ${web}"
        exit
    fi
    echo ${arr[@]}
}

backup() {
    file=$1
    cpfile="backup/${file}.txt"
    if [ -f $cpfile ]; then
        rm -f $cpfile
    fi
    if [ -f $file ]; then
        mv $file $cpfile
    fi
}

recover() {
    file=$1
    if [ -f $file ]; then
        rm -f $file
    fi
    if [ -f $cpfile ]; then
        mv "backup/${file}" $file
    fi
}

replace() {
    file=$1
    oldstr=$2
    newstr=$3
    # macOS环境中需要传入一个文件做备份，这里传入""就行
    sed -i "" "s/${oldstr}/${newstr}/g" $file
    if [ $? -ne 0 ]; then
        sed -i "s/${oldstr}/${newstr}/g" $file
    fi
    if [ $? -ne 0 ]; then
        recover $file
        echo "panic:replace error,file=${file},oldstr=${oldstr},newstr=${newstr}"
        exit
    fi
}

current_time=$(date "+%Y-%m-%d %H:%M:%S")
generator() {
    server_name=$1
    server_addr=$2
    filename="${server_name}.${domain}.conf.template"
    backup $filename
    cp ./script/conf.temp $filename
    replace $filename "__DOMAIN__" $domain
    replace $filename "__NAME__" $server_name
    replace $filename "__ADDR__" $server_addr
    replace $filename "__TIME__" $current_time
}

if [ ! -d "backup" ]; then
    mkdir "backup"
fi

for web in ${weblist[@]}; do
    arr=($(splitstr $web))
    generator ${arr[0]} ${arr[1]}
done

echo "----------- generator success -----------"
