#!/bin/bash

#只能在Linux下执行

#执行命令
#chmod a+x goinstall.sh
#./goinstall.sh

#选择要下载的版本
version="$1"
if [ -z $version ]; then
    echo "请输入要下载的go版本号"
    exit
fi

update_version="go${version}"

#备份当前脚本
backupshell() {
    basepath=$(
        cd $(dirname $0)
        pwd
    )
    #如果处于data中，那就不需要备份
    if [ $basepath = "/data" ]; then
        exit
    fi
    #如果存在，则备份
    if [ -f /data/goinstall.sh ]; then
        rm -f /data/goinstall-old.sh
        mv /data/goinstall.sh /data/goinstall-old.sh
    fi
    goinstallpath="${basepath}/goinstall.sh"
    cp $goinstallpath /data/goinstall.sh
}

if command -v go >/dev/null 2>&1; then
    version=$(go env GOVERSION)
    if [ $version = $update_version ]; then
        echo "---------- 当前go版本已经是 ${version} ，无需重复安装---------"
        backupshell
        exit
    fi
fi

echo "---------- 开始安装升级 ${update_version} ---------"

#判断/data目录是否存在，不存在就创建
if [ ! -d /data ]; then
    mkdir -p /data
fi

#整理链接，下载go包
filename="go${version}.linux-amd64.tar.gz"
filepath="/data/${filename}"
if [ ! -f $filepath ]; then
    gourl="http://mirrors.ustc.edu.cn/golang/${filename}"
    curl_code=$(curl -# -o $filepath --insecure -s -w %{http_code} $gourl)
    if [ $curl_code != "200" ]; then
        echo "下载失败: $gourl"
        rm -f $filepath
        exit
    fi
    echo "下载成功"
fi

#开始安装
#确保当前目录下没有解压的go包
rm -rf /data/go
#开始解压:
# z，通过gzip压缩或解压
# x，解压tar包
# f，指定压缩的文件名字
# C（大写），指定解压的目录路径
tar -zxf ${filepath} -C /data

#配置~/.bash_profile文件
configprofile() {
    #如果~/.bash_profile文件不存在
    profile=~/.bash_profile
    content="export PATH=\$PATH:${goroot}/bin"
    if [ ! -f $profile ]; then
        #创建~/.bash_profile文件
        echo $content >>$profile
        chmod a+x $profile
    else
        #同时做好防止重复配置
        if [ ! -f ~/.goenv_block ]; then
            echo $content >>$profile
            echo "防止重复配置~/.bash_profile" >>~/.goenv_block
        fi
    fi
}

#获取go root
#有可能不存在go command
#默认的 go root 路径
goroot=/usr/local/go
if command -v go >/dev/null 2>&1; then
    #系统原有的 go root 路径
    goroot=$(go env GOROOT)
else
    #go环境没有配置好
    $configprofile
fi

#移除旧的 go root 包
rm -rf $goroot
#添加新的 go root 包
mv /data/go $goroot

#只有go生效的时候，才会备份当前脚本
if ! command -v go >/dev/null 2>&1; then
    echo "---------- go 环境配置失败 ---------"
    echo "---------- go version 升级失败 ---------"
    exit
fi

backupshell

source ~/.bash_profile
version=$(go env GOVERSION)
if [ ! $version -eq $update_version ]; then
    echo "---------- ${version} -> ${update_version} 升级失败  ---------"
    exit
fi

echo "---------- 成功安装升级 ${version} ---------"
