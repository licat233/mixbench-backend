#!/bin/bash
###
 # @Author: licat
 # @Date: 2022-11-17 14:52:05
 # @LastEditors: licat
 # @LastEditTime: 2022-12-30 11:05:41
 # @Description: licat233@gmail.com
### 

# on mac
cd /opt/homebrew/opt/kafka/bin

#on docker
cd /opt/kafka/bin

#查看topic列表
kafka-topics --list --bootstrap-server localhost:9092
# on docker
kafka-topics.sh --list --bootstrap-server localhost:9092

#删除topic
kafka-topics --delete --topic goodluck-topic --bootstrap-server localhost:9092

#创建topic
kafka-topics --create --topic goodluck-topic --bootstrap-server localhost:9092
# on docker
kafka-topics.sh --create --topic test-topic --bootstrap-server localhost:9092

#查看某topic的详细
kafka-topics --describe --topic goodluck-topic --bootstrap-server localhost:9092

# ------

#查看consumer-group列表
kafka-consumer-groups --bootstrap-server localhost:9092 --list
# on docker 
kafka-consumer-groups.sh --bootstrap-server localhost:9092 --list

#删除consumer-group
kafka-consumer-groups --bootstrap-server localhost:9092 --delete --group luckdraw-group

#on docker
kafka-consumer-groups.sh --group luckdraw-group --bootstrap-server localhost:9092

# ------

#运行producer
kafka-console-producer --topic goodluck-topic --bootstrap-server localhost:9092
# on docker
kafka-console-producer.sh --topic mixbench-log --bootstrap-server localhost:9092

#运行consumer
kafka-console-consumer --topic goodluck-topic --from-beginning --bootstrap-server localhost:9092
# on docker
kafka-console-consumer.sh --topic mixbench-log --from-beginning --bootstrap-server localhost:9092