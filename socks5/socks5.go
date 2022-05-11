package socks5

import (
	"log"
	"net"
)

// Start server
func StartServer(config Config) {
	// set iface
	if config.Iface != "" {
		ief, err := net.InterfaceByName(config.Iface)
		if err != nil {
			log.Fatal("get Interface error", err)
		}
		addrs, err := ief.Addrs()
		if err != nil {
			log.Fatal(err)
		}
		config._outIface = ief
		config._outIP = addrs[0].(*net.IPNet).IP.To4()
		log.Printf("iface name: %v, out ip: %v", config.Iface, config._outIP.String())
	}
	// start udp server
	u := &UDPServer{config: config}
	udpConn := u.Start()
	// start tcp server
	t := &TCPServer{config: config, udpConn: udpConn}
	t.Start()
}
