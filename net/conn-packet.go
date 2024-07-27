package net

import (
	"net"
)

type packetConnection struct {
	connection
}

func (x *packetConnection) ReadFrom(b []byte) (int, net.Addr, error) {
	return x.real().ReadFrom(b)
}

func (x *packetConnection) WriteTo(b []byte, addr net.Addr) (int, error) {
	return x.real().WriteTo(b, addr)
}

func (x *packetConnection) real() net.PacketConn {
	return x.conn.(net.PacketConn)
}
