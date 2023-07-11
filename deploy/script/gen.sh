#!/bin/bash

#进入monitor mode
set -m

current_path=$(
    cd $(dirname $0) || exit 1
    pwd
)

# project_root="${current_path%%mixbench-backend*}mixbench-backend"
project_root="${current_path%%deploy*}"

nohup=$($project_root/genModd.sh);

# 创建build目录
if [ ! -d "${project_root}/build" ]; then
    mkdir "${project_root}/build"
fi

script_path=$current_path

cd $script_path || exit 1

gen_api() {
    name=$1
    micro=$2
    if [ "$micro" != "api" ]; then
        return
    fi
    docs_path="${project_root}/services/v1/${name}/docs"
    if [ ! -d "$docs_path" ]; then
        #如果目录不存在，则退出
        echo "提示: 目录${docs_path} 不存在"
        return
    fi
    cd $docs_path || exit 1
    filename="${name}.api"
    if [ ! -f $filename ]; then
        #如果文件不存在，则退出
        echo "提示: 文件${filename} 不存在"
        return
    fi
    nohup=$(goctl api go --api="$filename" --dir="../api" --style="goZero")
    if [ $? -ne 0 ]; then
        exit 1
    fi
    echo "已生成${name}-${micro}服务相关文件!"
    #生成接口文档swagger.json
    nohup=$(goctl api plugin -plugin goctl-swagger="swagger -filename ../api/static/swagger.json" -api "${name}.api" -dir .)
    if [ $? -ne 0 ]; then
        echo "开始生成${name}-${micro}的swagger.json文件失败！！"
        exit 1
    fi
    echo "已生成${name}-${micro}服务swagger.json文件"
    #生成typescript调用文件
    nohup=$(goctl api ts -api "${name}.api" -dir ../api/static)
    if [ $? -ne 0 ]; then
        echo "错误：生成ts文件失败"
        exit 1
    fi
    cd ../api/static || exit 1
    mdfile="${name}.ts"
    if [ ! -f $mdfile ]; then
        echo "提示: 文件 $mdfile 不存在!"
        return
    fi
    oldstr="* as components"
    newstr="type * as components"
    # macOS环境中需要传入一个文件做备份
    sed -i "" "s/${oldstr}/${newstr}/g" $mdfile
    if [ $? -ne 0 ]; then
        sed -i "s/${oldstr}/${newstr}/g" $mdfile
    fi
    if [ $? -ne 0 ]; then
        exit 1
    fi
    echo "已生成${name}-${micro}服务typescript接口文件"
    # cd ../
    # rm -f Dockerfile
    # goctl docker -go "${name}.go"
    # echo "生成${name} API服务的Dockerfile"
}

gen_rpc() {
    name=$1
    micro=$2
    if [ "$micro" != "rpc" ]; then
        return
    fi
    docs_path="${project_root}/services/v1/${name}/docs"
    if [ ! -d $docs_path ]; then
        #如果目录不存在，则退出
        echo "目录: ${docs_path} 不存在"
        return
    fi
    cd $docs_path
    filename="${name}.proto"
    if [ "${name}" == "auth" ]; then
        filename="authx.proto"
    fi
    if [ ! -f $filename ]; then
        #如果文件不存在，则退出
        echo -e "\n提示: 文件${filename} 不存在, 无需生成${name}-rpc服务"
        return
    fi
    # 判断该rpc服务是否为 multiple
    count=$(grep -c '^service.*{$' "$docs_path/$filename")
    if [ $count -gt 1 ]; then
        # 如果大于1，说明是 multiple
        nohup=$(goctl rpc protoc $filename --go_out=../rpc --go-grpc_out=../rpc --zrpc_out=../rpc --style=goZero --multiple=true)
    else
        nohup=$(goctl rpc protoc $filename --go_out=../rpc --go-grpc_out=../rpc --zrpc_out=../rpc --style=goZero)
    fi

    if [ $? -ne 0 ]; then
        exit 1
    fi
    pb_path="${project_root}/services/v1/${name}/rpc/${name}_pb"
    if [ ! -d $pb_path ]; then
        #如果目录不存在，则退出
        echo "目录: ${pb_path} 不存在!!!"
        return
    fi
    cd $pb_path

    # macOS环境中需要传入一个文件做备份
    sed -i "" 's/,omitempty//g' "${name}.pb.go"
    if [ $? -ne 0 ]; then
        sed -i 's/,omitempty//g' "${name}.pb.go"
    fi
    if [ $? -ne 0 ]; then
        exit 1
    fi

    echo -e "\n已生成${name}-${micro}服务相关文件!"
    # cd ../
    # rm -f Dockerfile
    # goctl docker -go "${name}.go"
    # echo "生成${name} RPC服务的Dockerfile"
}

gen_model() {
    name=$1
    sql_file="${project_root}/deploy/mysql/${name}.sql"
    if [ ! -f $sql_file ]; then
        #如果文件不存在，则退出
        echo -e "\n提示: 文件 ${name}.sql 不存在, 无需生成$name-model文件"
        return
    fi
    cd "${project_root}/services/v1/${name}/docs" || exit
    if ! goctl model mysql ddl --src "$sql_file" -dir ../model --cache --style goZero; then
        exit
    fi
}

source ./data.sh

#选择要构建的服务
service_name=$1
all_services=($(filter_service $service_name))
if [ -z $all_services ]; then
    echo "unknown service_name: ${service_name}"
    exit 1
fi

#选择要构建的子服务
micro_name=$2
all_micros=($(filter_micro $micro_name))
if [ -z $all_micros ]; then
    echo "unknown micro_name: ${micro_name}"
    exit 1
fi

has_sql_file() {
    name=$1
    sql_dir="${project_root}/deploy/mysql"
    sql_path="${sql_dir}/${name}.sql"
    if [ ! -f "$sql_path" ]; then
        return 1
    fi
    return 0
}

gen() {
    name=$1
    # ①使用 goctl 工具生产 model 文件
    # ②根据 genzero 配置文件，使用 genzero 工具生成相关的 .api和.proto 文件
    # ③根据 .api 和 .proto 文件，使用 goctl 工具生成相关的 gozero 微服务文件
    # ④根据 genzero 配置文件，使用 genzero 工具修改相关的 gozero 微服务文件
    # 注意：在某服务没有 mysql 文件的时候，不需要使用 genzero 生成该服务的 .api和.proto 文件
    docs_path="${project_root}/services/v1/${name}/docs"
    if [ ! -d $docs_path ]; then
        #如果目录不存在，则退出
        echo "目录: ${docs_path} 不存在"
        return
    fi

    # 判断是否存在 sql 文件， 如果存在才能使用 genzero 工具
    has_sql_file $name

    exist_sql_file=$? # 0 表示存在，1 表示不存在

    echo -e "\n开始生成${name}服务文件"



    cd "$docs_path" || exit
    # 使用 genzero
    if [ $exist_sql_file -eq 0 ]; then
        for micro in ${all_micros[*]}; do
            gen_model $name
        done
        # 如果不存在 genzero.yaml 文件，则生成
        if [ ! -f "./genzero.yaml" ]; then
            if ! genzero init config; then
                echo "${docs_path}目录下缺少 genzero.yaml 文件，已经生成，请先配置!"
                exit 1
            fi
        fi

        # ②根据yaml配置文件来生成微服务的配置文件
        if ! genzero start; then
            exit 1
        fi
    fi

    # ③使用 goctl 生成微服务文件
    for micro in ${all_micros[*]}; do
        gen_rpc $name $micro
        gen_api $name $micro
    done

    go mod tidy

    # 使用 genzero
    if [ $exist_sql_file -eq 0 ]; then
        # ④使用 genzero 修改微服务文件
        cd "$docs_path" || exit
        if ! genzero start; then
            exit 1
        fi
    fi

    echo -e "已生成${name}服务相关文件!"
}

#开始遍历编译
for name in ${all_services[*]}; do
    gen $name
done

echo "----------- generator successful -----------"
