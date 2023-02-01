#!/bin/bash

#进入monitor mode
set -m

current_path=$(
    cd $(dirname $0)
    pwd
)

cd $current_path

# 配置控制台访问令牌

# 授权kube-system默认服务账号
kubectl apply -f kube-system-default.yaml

# 对于Mac环境
TOKEN=$(kubectl -n kube-system describe secret default| awk '$1=="token:"{print $2}')
kubectl config set-credentials docker-desktop --token="${TOKEN}"
echo $TOKEN

# mac 查看token
$HOME/.kube/config

