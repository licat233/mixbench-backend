#!/bin/bash
###
# @Author: licat
# @Date: 2023-01-11 15:40:07
 # @LastEditors: licat
 # @LastEditTime: 2023-01-30 17:36:32
# @Description: licat233@gmail.com
###

#进入monitor mode
set -m

current_path=$(
    cd $(dirname $0)
    pwd
)

project_root="${current_path%%mixbench-backend*}mixbench-backend"

cd $project_root/deploy/sql2pb/build

# if [ -f mixbench.proto ]; then
#     rm -f mixbench.proto
# fi

# ./sql2pb -schema mixbench -package admin_proto -go_package admin_pb -service_name Admin >mixbench.proto

./sql2pb -schema mixbench_admin -package admin_proto -go_package admin_pb -service_name Admin -filename admin.proto -multiple=false
