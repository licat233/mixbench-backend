#!/bin/bash
###
 # @Author: licat
 # @Date: 2023-01-14 16:14:03
 # @LastEditors: licat
 # @LastEditTime: 2023-01-14 16:15:50
 # @Description: licat233@gmail.com
### 

dirname=./rpc

goctl rpc protoc addressbook.proto --go_out=$dirname --go-grpc_out=$dirname --zrpc_out=$dirname -style goZero

