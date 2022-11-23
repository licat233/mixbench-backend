#!/bin/bash
set -m
current_path=$(
    cd $(dirname $0)
    pwd
)
cd $current_path
pid=$(ps -ef | grep luckdraw-mq | grep -v grep | awk '{print $2}')
if [ ! -z $pid ]; then
    kill -9 $pid
fi
rm -f luckdraw-mq
go build -o luckdraw-mq luckdraw.go

#1.启动zookeeper服务
#二选一
#方式一，通过zookeeper/bin目录中的zkServer.sh文件启动
/usr/local/zookeeper/bin/zkServer.sh --config /usr/local/zookeeper/conf start
/usr/local/zookeeper/bin/zkServer.sh --config /usr/local/zookeeper/conf restart
#方式二，通过kafka/bin目录中的zookeeper-server-start.sh文件启动(注意改:dataDir=/tmp/zookeeper/kafka)
/usr/local/kafka/bin/zookeeper-server-start.sh /usr/local/kafka/config/zookeeper.properties
/usr/local/kafka/bin/zookeeper-server-stop.sh /usr/local/kafka/config/zookeeper.properties

#2.启动kafka服务，初次安装需要删除/tmp/kafka-logs
/usr/local/kafka/bin/kafka-server-start.sh /usr/local/kafka/config/server.properties
/usr/local/kafka/bin/kafka-server-stop.sh /usr/local/kafka/config/server.properties

#3.kafka创建一个主题(topic)
#使用 kafka-topics.sh 创建单分区单副本的 topic goodluck-topic
/usr/local/kafka/bin/kafka-topics.sh --bootstrap-server localhost:9092 --create --topic goodluck-topic

#查看topic列表
/usr/local/kafka/bin/kafka-topics.sh --bootstrap-server localhost:9092 --list

#consumer group
/usr/local/kafka/bin/kafka-consumer-groups.sh --bootstrap-server localhost:9092 --describe --all-groups

#查看消费者列表 --list
/usr/local/kafka/bin/kafka-consumer-groups.sh --bootstrap-server localhost:9092 --list