#!/bin/bash

#进入monitor mode
set -m

current_path=$(
    cd $(dirname $0)
    pwd
)

project_root="${current_path%%mixbench*}mixbench"
script_path="${project_root}/deploy/script/mac"
if [ ! -d $script_path ]; then
    #如果目录不存在，则退出
    echo "目录: ${script_path} 不存在"
    return
fi
cd $script_path

run() {
    filename=$1
    service_path=$2
    if [ ! -d $service_path ]; then
        #如果目录不存在，则退出
        # echo "目录: ${service_path} 不存在"
        return
    fi
    cd $service_path
    cd ../../../../build/$filename
    if [ ! -f $filename ]; then
        #如果文件不存在，则退出
        echo "文件: ${filename} 不存在"
        return
    fi
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
    #启动
    echo "后台启动 ${filename} 服务"
    rm -f "${service_path}/nohup.out"
    # nohup "${service_path}/${filename}" > ${filename}.file 2>&1 &
    nohup "${service_path}/${filename}" &
    return
    nohup=$("./${filename}" &)
    index=$(echo $(awk -v a="$nohup" -v b="Starting" 'BEGIN{print index(a,b)}'))
    if (($index == 1)); then
        echo "启动 ${filename} 失败: $nohup"
    else
        echo "启动 ${filename} 成功: $nohup"
    fi
}

#选择要构建的服务
service_name=$1
all_services=()
if [ -z $service_name ]; then
    all_services=("captcha" "admin" "luckdraw")
elif [[ "$service_name" = "captcha" ]] || [[ "$service_name" = "admin" ]] || [[ "$service_name" = "luckdraw" ]]; then
    all_services=($service_name)
else
    echo "unknown service_name: ${service_name}"
    exit
fi

#选择要构建的子服务
micro_name=$2
all_micros=()
if [ -z $micro_name ]; then
    all_micros=("rpc" "mq" "api")
elif [[ "$micro_name" = "rpc" ]] || [[ "$micro_name" = "mq" ]] || [[ "$micro_name" = "api" ]]; then
    all_micros=($micro_name)
else
    echo "unknown micro_name: ${micro_name}"
    exit
fi

#开始遍历编译
for name in ${all_services[*]}; do
    for micro in ${all_micros[*]}; do
        filename="${name}-${micro}"
        service_path="${project_root}/service/v1/${name}/${micro}"
        run $filename $service_path
    done
done

echo "----------- runing successful -----------"