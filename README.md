# mixbench-backend

## 系统结构组成

| 类别     | 服务名称      | 暴露端口 | 映射端口    | 说明                                             |
| -------- | ------------- | -------- | ----------- | ------------------------------------------------ |
| 数据库   | mysql         | 3306     | -           | mysql数据库                                      |
| 数据库   | mysqlmanage   | 80       | -           | mysql管理面板 (ui)                               |
| 缓存库   | redis         | 6379     | -           | redis缓存库                                      |
| 缓存库   | redismanage   | 80       | -           | redis管理面板(ui)                                |
| 公网网关 | nginx         | 80,443   | -           | nginx正向代理服务                                |
| 内网网关 | nginx         | 80,443   | 11080,11443 | nginx反向代理服务                                |
| 数据中心 | datacenter    | -        | -           | 数据卷容器，存储各个服务生产的数据(目前暂时弃用) |
| 文件管理 | filebrowser   | 80       | -           | 文件管理服务，用于浏览数据中心文件(ui)           |
| 链路追踪 | jaeger        | 6831/udp | -           | 接受jaeger.thriftThrift-compact 协议             |
| 链路追踪 | jaeger        | 6832/udp | -           | 接受jaeger.thriftThrift-binary 协议              |
| 链路追踪 | jaeger        | 5778     | -           | 服务配置（采样等）                               |
| 链路追踪 | jaeger        | 16686    | -           | 服务前端 (ui)                                    |
| 链路追踪 | jaeger        | 14268    | -           | 直接从客户端接受jaeger.thrift协议                |
| 链路追踪 | jaeger        | 9411     | -           | 兼容 Zipkin 服务（可选）                         |
| 服务监控 | prometheus    | 9090     | -           | 服务监控系统                                     |
| 服务监控 | grafana       | 3000     | -           | prometheus服务监控系统 (ui)                      |
| 数据收集 | elasticsearch | 9200     | -           | http协议连接es使用的端口                         |
| 数据收集 | elasticsearch | 9300     | -           | tcp客户端连接使用的端口                          |
| 数据收集 | kibana        | 5601     | -           | 查看elasticsearch数据 (ui)                       |
| 日志收集 | filebeat      | -        | -           | 收集日志并上报到kafka                            |
| 日志处理 | go-stash      | -        | -           | 处理kafka中的日志并发送到es                      |
| 消息队列 | zookeeper     | 2181     | -           | kafka的依赖服务                                  |
| 消息队列 | kafka         | 9092     | -           | kafka消息队列                                    |
| 消息队列 | CMAK          | 9000     | -           | 用于管理Apache Kafka集群的工具 (ui)              |
| 消息队列 | asynqmon      | 8080     | -           | go基于redis实现的消息队列及面板 (ui)             |
| 核心服务 | app           | -        | -           | modd热编译环境                                   |
| 核心业务 | admin-api     | 7011     | -           | 管理员API服务                                    |
| 核心业务 | admin-rpc     | 7012     | -           | 管理员RPC服务                                    |
| 核心业务 | captcha-api   | 7021     | -           | 验证码API服务                                    |
| 核心业务 | captcha-rpc   | 7022     | -           | 验证码RPC服务                                    |
| 核心业务 | luckydraw-api | 7031     | -           | 抽奖API服务                                      |
| 核心业务 | luckydraw-rpc | 7032     | -           | 抽奖RPC服务                                      |
| 核心业务 | luckydraw-mq  | 7033     | -           | 抽奖mq服务(kafka或asynq)                         |

## 说明

外部访问所有的服务均通过nginx反向代理: SEVER_NAME.mixbench.net
所有数据存储在datacenter数据卷容器中

## 注意事项

本地调试的时候，请关闭vpn

## 部署方式

docker-compose

>```shell
> docker-compose up -d
>```

## shell脚本管理服务使用说明

主要分为：生成(gen)、构建(build)、运行(run)、清理(clear)

入口文件: cmd.sh

使用说明:

```shell
./cmd.sh -h
```

command:

> gen [service-name] [micro-name]
> build [service-name] [micro-name]
> run [service-name] [micro-name]
> clear [service-name] [micro-name]
> [service-name] [micro-name]

## prometheus生态架构

基于prometheus的服务指标监控

![alt prometheus生态架构图](./deploy/docs/img/prometheus-struct.png)

## The simplified architecture that we use with [go-zero](https://github.com/zeromicro/go-zero)

![alt gozero](./deploy/docs/img/zo-zero-simplify.png)

## 关于日志处理方案([go-stash](https://github.com/kevwan/go-stash))

![alt go-stash](./deploy/docs/img/go-stash.png)

## goctl使用示例

① 生成 user.api 文件

>```shell
>goctl api -o user.api
>```

② 根据 user.api 文件生成 api 服务

>```shell
>goctl api go -api user.api -dir . -style goZero
>```

③ 启动 api 服务

>```shell
>go run user.go -f etc/user-api.yaml
>```

④ 根据 user.api 文件生成 swagger 文件

>确保已经安装了【 [goctl-swagger](https://github.com/zeromicro/goctl-swagger) 】插件:
>
>```shell
>GOPROXY=https://goproxy.cn/,direct go install github.com/zeromicro/goctl-swagger@latest
>```
>
>生成 user.json 文件:
>
>```shell
>goctl api plugin -plugin goctl-swagger="swagger -filename user.json" -api user.api -dir .
>```

## .gitignore的编写

忽略二进制文件

```shell
# 忽略掉所有文件
*
# 取消忽略所有带有后缀名的文件
!*.*
# 取消忽略所有目录
!*/
```

## 强制重置仓库

 .gitignore文件目录下CMD输入

```shell
git rm -r --cached .
```

 重置

```shell
git checkout --orphan latest_branch
git add -A
git commit -am newest
git branch -D main
git branch -m main
```

 提交

```shell
git push -f origin main
```

一键提交

```shell
git add .
git commit -m update --no-verify
git push origin main
git push origin-t main
```

## 放弃本地更改，强制覆盖本地

```shell
git fetch --all
git reset --hard origin/main
git pull origin main
```

注意：
1、 .gitignore只能忽略那些原来没有被track的文件，如果某些文件已经被纳入了版本管理中，则修改.gitignore是无效的。

2、 想要.gitignore起作用，必须要在这些文件不在暂存区中才可以，.gitignore文件只是忽略没有被staged(cached)文件， 对于已经被staged文件，加入ignore文件时一定要先从staged移除，才可以忽略。

## 生成客户端调用代码(go,javascript,php)

```shell
goctl api ts -api captcha.api -dir .
```

## .proto 和 .api 文件由 [sql2rpc](https://github.com/licat233/sql2rpc) 生成更新

`go install github.com/licat233/sql2rpc@latest`

```shell
sql2rpc -h
```
