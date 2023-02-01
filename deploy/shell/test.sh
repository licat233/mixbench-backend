#!/bin/bash

name=admin
pid=$(ps -ef | grep ${name}-rpc | grep -v grep | awk '{print $2}')
if [ ! -z $pid ]; then
    kill -9 $pid
    if [ ! $? -eq 0 ]; then
        #kill失败
        echo "${name}-rpc: [ kill -9 $pid ] 失败"
        exit
    fi
fi
service_path="/www/wwwroot/mixbench/service/v1/${name}/rpc"
cd $service_path

"./${name}-rpc" &
exit

name=admin
service_path="/www/wwwroot/mixbench/service/v1/${name}"
if [ ! -d $service_path ]; then
    echo "不存在"
else
    echo "存在"
fi

exit

output=$(ps -ef | grep admin-rpc | grep -v grep | awk '{print $2}')
echo $output
if [ ! -z $output ]; then
    echo "空"
fi

exit

loading_show=""
loading_index=0
loading_ch=('|' '\' '-' '/')
loading() {
    i=$loading_index
    j=$1
    if [ -z $j ]; then
        j=100
    fi
    str=$loading_show
    while [ $i -le $j ]; do
        let index=i%4
        if [ $i -le 20 ]; then
            let color=44
            let bg=34
        elif [ $i -le 45 ]; then
            let color=43
            let bg=33
        elif [ $i -le 75 ]; then
            let color=41
            let bg=31
        else
            let color=42
            let bg=32
        fi
        printf "\033[${color};${bg}m%-s\033[0m %d %c\r" "$str" "$i" "%" "${loading_ch[$index]}"
        usleep 30000
        let i=i+1
        if [ $(($i % 2)) -eq 0 ]; then
            str+="#"
        fi

    done
    loading_index=$i
    loading_show=$str
    if (($loading_index >= 100)); then
        printf "\n"
    fi
    unset str
}

loading 20
loading 50
loading 80
loading 100
exit

all_services=("captcha" "luckydraw" "admin")
for name in ${all_services[@]}; do
    echo $name
done

exit

echo $0
exit

fun() {
    echo "hello"
}
fun
exit

touch demo.txt
path=./demo.txt
rm -f $path
exit

path=~/.bash_profile
if [ -f $path ]; then
    echo "存在"
fi

exit

basepath=$(
    cd $(dirname $0)
    pwd
)
readonly basepath
echo $basepath
exit

chmod a+x goupdate.sh
/data/goupdate.sh 1.19

goroot="/usr/local/go"
profile="./.bash_profile"
content="export PATH=\$PATH:${goroot}/bin"
if [ ! -f "$profile" ]; then
    #创建~/.bash_profile文件
    touch $profile
    chmod a+x $profile
fi
# 在~/.bash_profile最后一行行后添加字符串
echo $content >>$profile

exit

goroot="default path"
if command -v go >/dev/null 2>&1; then
    goroot=$(go env GOROOT)
fi
echo $goroot
exit

version="$1"
if [ -z $version ]; then
    echo "请输入要下载的go版本号"
    exit
fi

curl_code=$(curl -# -O --insecure -s -w %{http_code} https://sscpic1.chinaz.net/Files/pic/pic9/202208/apic42541_s.jpg)
if [ $curl_code != "200" ]; then
    echo "下载失败"
    exit
fi
echo "下载成功"
