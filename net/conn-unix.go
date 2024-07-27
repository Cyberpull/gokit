package net

import (
	"net"
	"os"
	"syscall"
)

type UnixConn struct {
	packetConnection
}

func (x *UnixConn) File() (f *os.File, err error) {
	return x.real().File()
}

func (x *UnixConn) ReadFromUnix(b []byte) (int, *net.UnixAddr, error) {
	return x.real().ReadFromUnix(b)
}

func (x *UnixConn) ReadMsgUnix(b []byte, oob []byte) (n int, oobn int, flags int, addr *net.UnixAddr, err error) {
	return x.real().ReadMsgUnix(b, oob)
}

func (x *UnixConn) SetReadBuffer(bytes int) error {
	return x.real().SetReadBuffer(bytes)
}

func (x *UnixConn) SetWriteBuffer(bytes int) error {
	return x.real().SetWriteBuffer(bytes)
}

func (x *UnixConn) SyscallConn() (syscall.RawConn, error) {
	return x.real().SyscallConn()
}

func (x *UnixConn) WriteMsgUnix(b []byte, oob []byte, addr *net.UnixAddr) (n int, oobn int, err error) {
	return x.real().WriteMsgUnix(b, oob, addr)
}

func (x *UnixConn) WriteToUnix(b []byte, addr *net.UnixAddr) (int, error) {
	return x.real().WriteToUnix(b, addr)
}

func (x *UnixConn) CloseRead() error {
	return x.real().CloseRead()
}

func (x *UnixConn) CloseWrite() error {
	return x.real().CloseWrite()
}

func (x *UnixConn) real() *net.UnixConn {
	return x.conn.(*net.UnixConn)
}

// ===========================

func newUnixConn(conn *net.UnixConn) *UnixConn {
	value := &UnixConn{}
	initConn(&value.connection, conn)
	return value
}
