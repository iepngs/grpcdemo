# grpcdemo
gRPC with Go and PHP client


0. 请预先准备好gprc的环境

1. 运行server

```shell
$ go run server.go
```

2. 运行grpc的go客户端

```shell
$ go run client/go/main.go
```
3. 运行grpc的php客户端
运行php客户端需要composer初始化一次项目:

```shell
$ cd client/php
$ composer update
$ composer dump-autoload
# 然后运行:
$ php index.php
```
