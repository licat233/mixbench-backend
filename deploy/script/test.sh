#!/bin/bash

#进入monitor mode
set -m

current_path=$(
    cd $(dirname $0)
    pwd
)

cd $current_path

# source "./utils.sh"



# project_root="${current_path%%mixbench-backend*}mixbench-backend"
project_root="${current_path%%deploy*}"
# echo $project_root

services_path="${project_root}/services/v1"

cd $services_path

count=$(grep -c '^service.*{$' "$services_path/luckydraw/docs/luckydraw.proto")

echo $count

exit 0

# services=$(ls -d *)

# echo $services

# if [[ "a" != "a" ]]; then
#     echo 233
# fi
# echo over
