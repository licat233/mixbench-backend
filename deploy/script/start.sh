#!/bin/bash
set -m

#判断brew是否存在
if false && command -v brew >/dev/null 2>&1; then
    #启动redis
    brew services start redis

    #启动etcd
    brew services start etcd

    #启动kafka
    brew services start kafka

    #启动nginx
    brew services start nginx
fi

current_path=$(
    cd $(dirname $0)
    pwd
)

# 从 string 字符串最后一次出现 *chars 的位置开始，截取 *chars 左边的所有字符
# project_root="${current_path%%mixbench-backend*}mixbench-backend" 
# script_path="${project_root}/deploy/script/mac"

script_path=$current_path

if [ ! -d $script_path ]; then
    #如果目录不存在，则退出
    echo "目录: ${script_path} 不存在"
    return
fi
cd $script_path

if [[ "$1" = "gen" ]];then
./gen.sh $2 $3
elif [[ "$1" = "build" ]];then
./build.sh $2 $3
elif [[ "$1" = "run" ]];then
./run.sh $2 $3
elif [[ "$1" = "clear" ]];then
./run.sh $2 $3
elif [[ "$1" = "-h" ]];then
echo "command: "
echo "gen [service-name] [micro-name]"
echo "build [service-name] [micro-name]"
echo "run [service-name] [micro-name]"
echo "clear [service-name] [micro-name]"
echo "[service-name] [micro-name]"
else
#开始 gen
nohup=$(./gen.sh $1 $2) 
#开始 编译
# nohup=$(./build.sh $1 $2) # 交由modd来编译
#开始 运行
# nohup=$(./run.sh $1 $2)  # 交由modd来运行
fi
