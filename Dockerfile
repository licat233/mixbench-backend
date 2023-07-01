FROM golang:alpine AS builder

#搭建开发环境
LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata bash git
ENV TZ Asia/Shanghai

WORKDIR /go

RUN go install github.com/zeromicro/go-zero/tools/goctl@latest
RUN go install github.com/zeromicro/goctl-swagger@latest
# goctl一键检查安装所需插件
RUN goctl env check -i -f --verbose
# modd热编译工具
RUN go install github.com/cortesi/modd/cmd/modd@latest

CMD [ "modd" ]
