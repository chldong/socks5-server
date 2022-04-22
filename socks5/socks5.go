package socks5

// Start server
func StartServer(config Config) {
	// start udp server
	u := &UDPServer{config: config}
	udpConn := u.Start()
	// start tcp server
	t := &TCPServer{config: config, udpConn: udpConn}
	t.Start()
}
