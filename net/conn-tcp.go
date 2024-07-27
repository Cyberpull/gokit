package net

import (
	"net"
	"os"
)

type TCPConn struct {
	packetConnection
}

func (x *TCPConn) File() (f *os.File, err error) {
	return x.real().File()
}

func (x *TCPConn) MultipathTCP() (bool, error) {
	return x.real().MultipathTCP()
}

func (x *TCPConn) SetReadBuffer(bytes int) error {
	return x.real().SetReadBuffer(bytes)
}

func (x *TCPConn) SetWriteBuffer(bytes int) error {
	return x.real().SetWriteBuffer(bytes)
}

func (x *TCPConn) CloseRead() error {
	return x.real().CloseRead()
}

func (x *TCPConn) CloseWrite() error {
	return x.real().CloseWrite()
}

func (x *TCPConn) real() *net.TCPConn {
	return x.conn.(*net.TCPConn)
}

// ===========================

func newTCPConn(conn *net.TCPConn) *TCPConn {
	value := &TCPConn{}
	initConn(&value.connection, conn)
	return value
}
