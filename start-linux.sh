#!/bin/bash

set -m

current_path=$(
    cd $(dirname $0)
    pwd
)

cd "$current_path/deploy/script/linux"

chmod a+x ./gen.sh
chmod a+x ./build.sh
chmod a+x ./run.sh
chmod a+x ./start.sh

./start.sh $1 $2 $3
