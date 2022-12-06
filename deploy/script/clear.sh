#!/bin/bash

#进入monitor mode
set -m

current_path=$(
    cd $(dirname $0)
    pwd
)

project_root="${current_path%%deploy*}"

script_path=$current_path

cd $script_path

clear_build(){
    name=$1
    micro=$2
    filename="${name}-${micro}"
    if [[ -z $name ]] || [[ -z $micro ]]; then
        echo "未输入build参数"
        exit
    fi

    service_path="${project_root}/service/v1/${name}/${micro}"
    if [ ! -d $service_path ]; then
        #如果不存在，说明没有这个服务，直接返回
        return
    fi
    #存在就进入
    cd $service_path
    #查找相关服务，并终止
    pid=$(ps -ef | grep $filename | grep -v grep | awk '{print $2}')
    if [ ! -z $pid ]; then
        kill -9 $pid
        if [ ! $? -eq 0 ]; then
            #kill失败
            echo "提示: ${filename}: [ kill -9 $pid ] 失败"
            exit
            return
        fi
        echo "已终止进程：${filename} $pid"
    fi
    #清除日志文件
    rm -rf ./logs/*
    if [ ! -f $filename ]; then
        #如果文件不存在，则退出
        echo "文件: ${filename} 不存在"
        return
    fi
    #清除二进制文件
    rm -f $filename
    
    echo "${filename} 已清除"
}

source ./data.sh

#选择要构建的服务
service_name=$1
all_services=($(filter_service $service_name))
if [ -z $all_services ]; then
    echo "unknown service_name: ${service_name}"
    exit
fi

#选择要构建的子服务
micro_name=$2
all_micros=($(filter_micro $micro_name))
if [ -z $all_micros ]; then
    echo "unknown micro_name: ${micro_name}"
    exit
fi

#开始遍历编译
for name in ${all_services[*]}; do
    for micro in ${all_micros[*]}; do
        clear_build $name $micro
    done
done