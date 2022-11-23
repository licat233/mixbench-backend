#!/bin/bash

#进入monitor mode
set -m

current_path=$(
    cd $(dirname $0)
    pwd
)

project_root="${current_path%%mixbench*}mixbench"
script_path="${project_root}/deploy/script"
if [ ! -d $script_path ]; then
    #如果目录不存在，则退出
    echo "目录: ${script_path} 不存在"
    return
fi
cd $script_path

gen_api() {
    name=$1
    micro=$2
    if [ "$micro" != "api" ]; then
        return
    fi
    service_path="${project_root}/service/v1/${name}/api/docs"
    if [ ! -d $service_path ]; then
        #如果目录不存在，则退出
        echo "目录: ${service_path} 不存在"
        return
    fi
    cd $service_path
    filename="${name}.api"
    if [ ! -f $filename ]; then
        #如果文件不存在，则退出
        echo "文件: ${filename} 不存在"
        return
    fi
    echo -e "\n开始生成${name}-${micro}服务..."
    nohup=$(goctl api go -api $filename -dir=../ -style goZero)
    echo "${name}-${micro}服务生成完成!"
    #生成接口文档swagger.json
    echo "生成接口文档swagger.json ...... "
    goctl api plugin -plugin goctl-swagger="swagger -filename ../static/swagger.json" -api "${name}.api" -dir .
    #生成typescript调用文件
    echo "生成typescript调用文件 ...... "
    goctl api ts -api "${name}.api" -dir ../static
}

gen_rpc() {
    name=$1
    micro=$2
    if [ "$micro" != "rpc" ]; then
        return
    fi
    service_path="${project_root}/service/v1/${name}/rpc/docs"
    if [ ! -d $service_path ]; then
        #如果目录不存在，则退出
        echo "目录: ${service_path} 不存在"
        return
    fi
    cd $service_path
    filename="${name}.proto"
    if [ ! -f $filename ]; then
        #如果文件不存在，则退出
        echo "文件: ${filename} 不存在"
        return
    fi
    echo -e "\n开始生成${name}-${micro}服务..."
    nohup=$(goctl rpc protoc $filename --go_out=../ --go-grpc_out=../ --zrpc_out=../ -style goZero)
    pb_path="${project_root}/service/v1/${name}/rpc/${name}_pb"
    if [ ! -d $pb_path ]; then
        #如果目录不存在，则退出
        echo "目录: ${pb_path} 不存在!!!"
        return
    fi
    cd $pb_path
    sed -i "" 's/,omitempty//g' "${name}.pb.go"
    echo "${name}-${micro}服务生成完成!"
}

run() {
    filename=$1
    service_path=$2
    if [ ! -d $service_path ]; then
        #如果目录不存在，则退出
        echo "目录: ${service_path} 不存在"
        return
    fi
    cd $service_path
    if [ ! -f $filename ]; then
        #如果文件不存在，则退出
        echo "文件: ${filename} 不存在"
        return
    fi
    echo "后台启动 ${filename} 服务"
    rm -f "${service_path}/nohup.out"
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

build_go() {
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

    echo -e "\n开始编译${filename} ...... "
    rm -f $filename
    go build -o $filename "${name}.go"
    echo "${filename} 编译完成"
    run $filename $service_path
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
        gen_rpc $name $micro
        gen_api $name $micro
        build_go $name $micro
    done
done

echo "----------- Initialize successful -----------"
