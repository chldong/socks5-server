package socks5

// Config is the socks5 server config
type Config struct {
	LocalAddr string
	Username  string
	Password  string
	ServerKey string
	ServerPem string
	TLS       bool
	Iface     string
}
