#!/bin/bash

current_path=$(
    cd $(dirname $0)
    pwd
)

cd $current_path

name="luckydraw"

# if [ "$1" != "ddl" ]; then
#     goctl model mysql datasource -url="mixbench:@tcp(127.0.0.1:3306)/mixbench_${name}" -table="*" -dir='.' -cache=true -style goZero
#     exit
# fi

project_root="${current_path%%mixbench-backend*}mixbench-backend"
sql_file="${project_root}/deploy/mysql/${name}.sql"
if ! goctl model mysql ddl --src "$sql_file" -dir . --cache --style goZero; then
    exit
fi

sql2rpc -model -db_schema="mixbench_${name}" -db_table="*" -dir="."
