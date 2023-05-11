#!/bin/bash
###
# @Author: licat
# @Date: 2023-01-15 14:05:33
# @LastEditors: licat
# @LastEditTime: 2023-02-20 18:27:20
# @Description: licat233@gmail.com
###

#进入monitor mode
set -m

current_path=$(
    cd $(dirname $0) || exit
    pwd
)

cd $current_path || exit

# project_root="${current_path%%mixbench-backend*}mixbench-backend"

name="luckydraw"
filename="luckydraw.api"

if ! sql2rpc -api -db_schema=mixbench_${name} -filename=${filename} -service_name=${name}-api -api_jwt="Auth" -api_middleware="AdminAuthMiddleware" -api_prefix="/v1/api/admin/luckydraw" --api_multiple=true; then
    exit 1
fi

goctl api go -api $filename -dir=../ -style goZero
