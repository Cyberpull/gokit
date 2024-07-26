package net

import (
	"net"
	"os"
	"syscall"
)

type IPConn struct {
	packetConnection
}

func (x *IPConn) File() (f *os.File, err error) {
	return x.real().File()
}

func (x *IPConn) ReadFromIP(b []byte) (int, *net.IPAddr, error) {
	return x.real().ReadFromIP(b)
}

func (x *IPConn) ReadMsgIP(b []byte, oob []byte) (n int, oobn int, flags int, addr *net.IPAddr, err error) {
	return x.real().ReadMsgIP(b, oob)
}

func (x *IPConn) SetReadBuffer(bytes int) error {
	return x.real().SetReadBuffer(bytes)
}

func (x *IPConn) SetWriteBuffer(bytes int) error {
	return x.real().SetWriteBuffer(bytes)
}

func (x *IPConn) SyscallConn() (syscall.RawConn, error) {
	return x.real().SyscallConn()
}

func (x *IPConn) WriteMsgIP(b []byte, oob []byte, addr *net.IPAddr) (n int, oobn int, err error) {
	return x.real().WriteMsgIP(b, oob, addr)
}

func (x *IPConn) WriteToIP(b []byte, addr *net.IPAddr) (int, error) {
	return x.real().WriteToIP(b, addr)
}

func (x *IPConn) real() *net.IPConn {
	return x.packetConnection.conn.(*net.IPConn)
}

// ===========================

func newIPConn(conn *net.IPConn) *IPConn {
	value := &IPConn{}
	initConn(&value.connection, conn)
	return value
}
