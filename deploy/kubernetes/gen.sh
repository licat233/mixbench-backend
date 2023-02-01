#!/bin/bash

#进入monitor mode
set -m

current_path=$(
    cd $(dirname $0)
    pwd
)

cd $current_path

# goctl kube deploy --name mixbench --image greet:v1 --namespace default --port 8888 --nodePort 32000 -o greet.yaml
goctl kube deploy --name mixbench-k8s --image mixbench-k8s:v1 --namespace default --port 80 --nodePort 11080 -o mixbench-k8s.yaml

