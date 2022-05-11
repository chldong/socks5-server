package socks5

import "net"

type Config struct {
	LocalAddr string
	Username  string
	Password  string
	ServerKey string
	ServerPem string
	TLS       bool
	Iface     string
	_outIP    net.IP
	_outIface *net.Interface
}
