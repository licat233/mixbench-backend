#!/bin/bash

#进入monitor mode
set -m

current_path=$(
    cd $(dirname $0)
    pwd
)

cd $current_path

# 安装
kubectl apply -f kubernetes-dashboard.yaml

# 检查 kubernetes-dashboard 应用状态
kubectl get pod -n kubernetes-dashboard

# 开启 API Server 访问代理
kubectl proxy

# 通过如下 URL 访问 Kubernetes dashboard
http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/



