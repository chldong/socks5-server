# socks5-server

A socks5 server written in Golang.

[EN](https://github.com/net-byte/socks5-server/blob/master/README.md) | [中文](https://github.com/net-byte/socks5-server/blob/master/README_CN.md)

[![Travis](https://travis-ci.com/net-byte/socks5-server.svg?branch=main)](https://github.com/net-byte/socks5-server)
[![Go Report Card](https://goreportcard.com/badge/github.com/net-byte/socks5-server)](https://goreportcard.com/report/github.com/net-byte/socks5-server)
![image](https://img.shields.io/badge/License-MIT-orange)
![image](https://img.shields.io/badge/License-Anti--996-red)
![image](https://img.shields.io/github/downloads/net-byte/socks5-server/total.svg)

# Features
* Support connect
* Support udp associate
* Support tcp over tls
* Support specified interface

# Usage
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
  -iface string
        specified interface
```

# Docker
[docker image](https://hub.docker.com/r/netbyte/socks5-server)

## Run server
```
docker run  -d --restart=always --net=host \
-p 1080:1080 -p 1080:1080/udp --name socks5-server netbyte/socks5-server -l :1080
```

## Run server with auth
```
docker run  -d --restart=always --net=host \
-p 1080:1080 -p 1080:1080/udp --name socks5-server netbyte/socks5-server -l :1080 -u root -p 123456
```

## Run server over tls with auth
```
docker run  -d --restart=always --net=host \
-p 1080:1080 -p 1080:1080/udp --name socks5-server netbyte/socks5-server -l :1080 -u root -p 123456 -tls -sk /app/certs/server.key -sp /app/certs/server.pem
```

## Run server with specified interface
```
docker run  -d --restart=always --net=host \
-p 1080:1080 -p 1080:1080/udp --name socks5-server netbyte/socks5-server -l :1080 -iface tun0
```

# License
[The MIT License (MIT)](https://raw.githubusercontent.com/net-byte/socks5-server/main/LICENSE)
