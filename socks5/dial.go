package socks5

import (
	"net"
	"time"

	"github.com/net-byte/socks5-server/sockopt"
)

// Dial is a helper function to dial a tcp or udp connection
func dial(network, addr string, outIface *net.Interface, outIP net.IP) (net.Conn, error) {
	var la net.Addr
	switch network {
	case "tcp":
		la = &net.TCPAddr{IP: outIP}
	case "udp":
		la = &net.UDPAddr{IP: outIP}
	}

	dialer := &net.Dialer{LocalAddr: la, Timeout: time.Duration(Timeout) * time.Second}
	if outIface != nil {
		dialer.Control = sockopt.Control(sockopt.Bind(outIface))
	}

	c, err := dialer.Dial(network, addr)
	if err != nil {
		return nil, err
	}

	if c, ok := c.(*net.TCPConn); ok {
		c.SetKeepAlive(true)
	}

	c.SetDeadline(time.Now().Add(time.Duration(Timeout) * time.Second))
	return c, err
}
