# socks5-server

一款基于Go语言开发的Socks5代理

[EN](https://github.com/net-byte/socks5-server/blob/master/README.md) | [中文](https://github.com/net-byte/socks5-server/blob/master/README_CN.md)

[![Travis](https://travis-ci.com/net-byte/socks5-server.svg?branch=main)](https://github.com/net-byte/socks5-server)
[![Go Report Card](https://goreportcard.com/badge/github.com/net-byte/socks5-server)](https://goreportcard.com/report/github.com/net-byte/socks5-server)
![image](https://img.shields.io/badge/License-MIT-orange)
![image](https://img.shields.io/badge/License-Anti--996-red)

# 特性
* 支持 connect
* 支持 udp associate
* 支持 tls

# 用法
```
Usage of /main:
  -l string
        local address (default ":1080")
  -p string
        password
  -u string
        username
  -sk string
        server key file path (default "../certs/server.key")
  -sp string
        server pem file path (default "../certs/server.pem")
  -tls enable tls
```

# Docker
[镜像](https://hub.docker.com/r/netbyte/socks5-server)

## 运行服务
```
docker run  -d --restart=always --net=host \
-p 1080:1080 -p 1080:1080/udp --name socks5-server netbyte/socks5-server -l :1080
```

## 运行服务（授权模式）
```
docker run  -d --restart=always --net=host \
-p 1080:1080 -p 1080:1080/udp --name socks5-server netbyte/socks5-server -l :1080 -u root -p 123456
```

## 运行服务（TLS授权模式）
```
docker run  -d --restart=always --net=host \
-p 1080:1080 -p 1080:1080/udp --name socks5-server netbyte/socks5-server -l :1080 -u root -p 123456 -tls -sk /app/certs/server.key -sp /app/certs/server.pem
```

# License
[The MIT License (MIT)](https://raw.githubusercontent.com/net-byte/socks5-server/main/LICENSE)
