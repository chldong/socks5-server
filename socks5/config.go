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
	_outIP    string
	_outIface *net.Interface
}
