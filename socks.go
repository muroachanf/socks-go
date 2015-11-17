package socks

import (
	"io"
	"log"
	"net"
)

const (
	socks4Version  = 0x04
	socks5Version  = 0x05
	cmdConnect     = 0x01
	addrTypeIPv4   = 0x01
	addrTypeDomain = 0x03
	addrTypeIPv6   = 0x04
)

type SocksConn struct {
	ClientConn net.Conn
}

func (s *SocksConn) Serve() {
	buf := make([]byte, 1)

	// read version
	io.ReadFull(s.ClientConn, buf)

	switch buf[0] {
	case socks4Version:
		s4 := socks4Conn{client_conn: s.ClientConn}
		s4.Serve()
	case socks5Version:
		s5 := socks5Conn{client_conn: s.ClientConn}
		s5.Serve()
	default:
		log.Printf("error version %s", buf[0])
		s.ClientConn.Close()
	}
}