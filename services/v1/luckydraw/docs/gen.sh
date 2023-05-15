#!/bin/bash

#进入monitor mode
set -m

current_path=$(
    cd "$(dirname $0)" || exit
    pwd
)

cd "$current_path" || exit

nohup=$(./goctl.sh)
./genzero.sh
