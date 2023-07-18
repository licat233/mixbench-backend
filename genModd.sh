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

base_path="services/v1"
common_path="common/**/*.go"
build_dir="build"
echo "monitored services:"
for service in ${service_list[*]}; do
    service_path="${base_path}/${service}"
    if [[ ! -d $service_path ]]; then
        continue
    fi
    echo ">${service}"
    echo "# ======================== ${service} start ========================" >>$filename
    for micro in ${micro_list[*]}; do
        micro_path="${service_path}/${micro}"
        if [[ ! -d $micro_path ]]; then
            continue
        fi
        if [[ ! -f "${micro_path}/${service}.go" ]]; then
            continue
        fi
        echo -e "\t-${micro}"
        build_file="${service}-${micro}"
        build_path="${build_dir}/${build_file}/${build_file}"
        go_path="${micro_path}/${service}.go"
        # yaml_file="${build_dir}/${build_file}/etc/${service}.yaml"
        yaml_file="${micro_path}/etc/${service}.yaml"
        # monitor_file="${micro_path}/**/*.go ${service_path}/model/**/*.go common/**/*.go"
        monitor_file="${micro_path}/**/*.go ${service_path}/model/**/*.go"
        prep="    prep: go build -o ${build_path} -v ${go_path}"
        # daemon="    daemon +sigkill: ${build_path} -f ${yaml_file}"
        daemon="    daemon +sigkill: ${build_path} -f ${yaml_file}"
        # if [[ "$micro" = "api" ]]; then
        #     daemon="${daemon} -d ${micro_path}"
        # fi
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

echo -e "\n------------------- successfully configured modd.conf -------------------"
