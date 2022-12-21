#!/bin/sh

docker run \
    -p 9090:9090 \
    -v /Users/licat/Desktop/project/mixbench/mixbench-backend/deploy/prometheus/server/prometheus.yml:/etc/prometheus/prometheus.yml \
    prom/prometheus:main