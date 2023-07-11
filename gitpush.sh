#!/bin/bash

#进入monitor mode
set -m

current_path=$(
    cd $(dirname $0)
    pwd
)

cd $current_path

COMMENT="$1"

if [ "$COMMENT" == "" ] || [ "$COMMENT" == "," ]; then
    COMMENT="captcha -> authentication"
fi

git add .
git commit -m "$COMMENT" --no-verify
git push origin main

echo "git push 已完成"
