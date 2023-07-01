#!/bin/bash

#进入monitor mode
set -m

current_path=$(
    cd $(dirname $0)
    pwd
)

cd $current_path

# project_root="${current_path%%mixbench-backend*}mixbench-backend"
project_root="${current_path%%deploy*}"
# echo $project_root

services_path="${project_root}/services/v1"

cd $services_path

services=$(ls -d *)

echo $services

# if [[ "a" != "a" ]]; then
#     echo 233
# fi
# echo over
