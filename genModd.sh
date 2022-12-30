#!/bin/bash
###
 # @Author: licat
 # @Date: 2022-12-06 14:38:15
 # @LastEditors: licat
 # @LastEditTime: 2022-12-30 10:01:21
 # @Description: licat233@gmail.com
### 

#进入monitor mode
set -m

current_path=$(
    cd $(dirname $0)
    pwd
)

cd $current_path

source "${current_path}/deploy/script/data.sh"

filename="modd.conf"
current_time=$(date "+%Y-%m-%d %H:%M:%S")
if [ -f $filename ]; then
    if [ -f "${filename}.old" ]; then
        rm -f "${filename}.old"
    fi
    mv $filename "${filename}.old"
    echo "# backup time: $current_time" >>"${filename}.old"
fi

touch $filename

# service/v1/admin/rpc/**/*.go {
#     prep: go build -o build/admin-rpc/admin-rpc -v service/v1/admin/rpc/admin.go
#     daemon +sigkill: build/admin-rpc/admin-rpc -f service/v1/admin/rpc/etc/admin.yaml
# }
# service/v1/admin/api/**/*.go {
#     prep: go build -o build/admin-api/admin-api -v service/v1/admin/api/admin.go
#     daemon +sigkill: build/admin-api/admin-api -f service/v1/admin/api/etc/admin-api.yaml -root=service/v1/admin/api
# }

base_path="service/v1"
common_path="common/**/*.go"
build_dir="build"
for service in ${service_list[*]}; do
    service_path="${base_path}/${service}"
    if [[ ! -d $service_path ]]; then
        continue
    fi
    echo "# ======================== ${service} start ========================" >>$filename
    for micro in ${micro_list[*]}; do
        micro_path="${service_path}/${micro}"
        if [[ ! -d $micro_path ]]; then
            continue
        fi
        build_file="${service}-${micro}"
        build_path="${build_dir}/${build_file}/${build_file}"
        go_path="${micro_path}/${service}.go"
        yaml_name="${service}"
        if [[ "$micro" != "rpc" ]]; then
            yaml_name="${service}-${micro}"
        fi
        yaml_path="${micro_path}/etc/${yaml_name}.yaml"
        monitor_file="${micro_path}/**/*.go common/**/*.go"
        prep="    prep: go build -o ${build_path} -v ${go_path}"
        daemon="    daemon +sigkill: ${build_path} -f ${yaml_path}"
        if [[ "$micro" = "api" ]]; then
            daemon="${daemon} -root=${micro_path}"
        fi
        echo "# ${micro}
${monitor_file} {
$prep
$daemon
}
" >>$filename
    done
    echo "# ======================== ${service} end ========================" >>$filename
    echo >>$filename
done

echo "------------------- success -------------------"
