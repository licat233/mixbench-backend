#!/bin/bash

service_list=("admin" "captcha" "luckdraw")

micro_list=("rpc" "api" "mq")

filter() {
    type=$1
    list=()
    case $type in
    "service")
        list=(${service_list[*]})
        ;;
    "micro")
        list=(${micro_list[*]})
        ;;
    *)
        echo ""
        return
        ;;
    esac

    name=$2
    if [ -z $name ]; then
        echo ${list[*]}
        return
    fi

    for item in ${list[*]}; do
        if [[ "$name"x = "$item"x ]]; then
            echo $name
            return
        fi
    done
    echo ""
}

# 如果返回 ""，则表示 unknown service_name
filter_service() {
    service_name=$1
    type=service
    res=($(filter $type $service_name))
    echo ${res[*]}
}

# 如果返回 ""，则表示 unknown micro_name
filter_micro() {
    micro_name=$1
    type=micro
    res=($(filter $type $micro_name))
    echo ${res[*]}
}

# list=($(filter_micro mq))
# echo ${list[*]}
