#!/bin/bash
###
 # @Author: licat
 # @Date: 2022-11-17 14:55:33
 # @LastEditors: licat
 # @LastEditTime: 2023-01-27 16:24:28
 # @Description: licat233@gmail.com，输入ddl，则执行本地文件
### 

current_path=$(
    cd $(dirname $0)
    pwd
)

cd $current_path

if [ "$1" != "ddl" ]; then
goctl model mysql datasource -url="mixbench:123123123@tcp(127.0.0.1:3306)/mixbench_admin" -table="*"  -dir='.' -cache=true -style goZero
exit
fi

project_root="${current_path%%mixbench-backend*}mixbench-backend"
sql_file="${project_root}/deploy/mysql/admin.sql"
goctl model mysql ddl --src "$sql_file" -dir . --cache --style goZero
