#!/bin/bash

#进入monitor mode
set -m

current_path=$(
    cd $(dirname $0)
    pwd
)

cd $current_path

if [[ "a" != "a" ]]; then
    echo 233
fi
echo over
