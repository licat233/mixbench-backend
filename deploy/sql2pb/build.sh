#!/bin/bash
###
# @Author: licat
# @Date: 2023-01-11 15:40:07
 # @LastEditors: licat
 # @LastEditTime: 2023-01-15 14:09:16
# @Description: licat233@gmail.com
###

if [ -f build/sql2pb ]; then
    rm -f build/sql2pb
fi

nohup=$(go build -o build/sql2pb sql2pb.go)
