FROM golang:1.19.3-alpine AS builder

#主要搭建modd环境
LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata bash git
ENV TZ Asia/Shanghai

WORKDIR /go

RUN go install github.com/cortesi/modd/cmd/modd@latest

CMD [ "modd" ]
