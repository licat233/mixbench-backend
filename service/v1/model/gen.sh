#!/bin/bash

if [ "$1" != "local"]; then
goctl model mysql datasource -url="mixbench:123123123@tcp(127.0.0.1:3306)/mixbench" -table="*"  -dir='.' -cache=true -style goZero
exit
fi

current_path=$(
    cd $(dirname $0)
    pwd
)

project_root="${current_path%%mixbench*}mixbench"
mysql_name="${project_root}/deploy/mysql/mixbench.sql"
goctl model mysql ddl -src="$mysql_name" -dir="." -cache=true -style goZero
