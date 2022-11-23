#!/bin/bash

cd /opt/homebrew/opt/kafka/bin

#查看topic列表
kafka-topics --list --bootstrap-server localhost:9092

#删除topic
kafka-topics --delete --topic goodluck-topic --bootstrap-server localhost:9092

#创建topic
kafka-topics --create --topic goodluck-topic --bootstrap-server localhost:9092

#查看某topic的详细
kafka-topics --describe --topic goodluck-topic --bootstrap-server localhost:9092

# ------

#查看consumer-group列表
kafka-consumer-groups --bootstrap-server localhost:9092 --list

#删除consumer-group
kafka-consumer-groups --bootstrap-server localhost:9092 --delete --group luckdraw-group

# ------

#运行producer
kafka-console-producer --topic goodluck-topic --bootstrap-server localhost:9092

#运行consumer
kafka-console-consumer --topic goodluck-topic --from-beginning --bootstrap-server localhost:9092
