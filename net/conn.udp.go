package net

import (
	"net"
	"net/netip"
	"os"
	"syscall"
)

type UDPConn struct {
	packetConnection
}

func (x *UDPConn) File() (f *os.File, err error) {
	return x.real().File()
}

func (x *UDPConn) ReadFromUDP(b []byte) (n int, addr *net.UDPAddr, err error) {
	return x.real().ReadFromUDP(b)
}

func (x *UDPConn) ReadFromUDPAddrPort(b []byte) (n int, addr netip.AddrPort, err error) {
	return x.real().ReadFromUDPAddrPort(b)
}

func (x *UDPConn) ReadMsgUDP(b []byte, oob []byte) (n int, oobn int, flags int, addr *net.UDPAddr, err error) {
	return x.real().ReadMsgUDP(b, oob)
}

func (x *UDPConn) ReadMsgUDPAddrPort(b []byte, oob []byte) (n int, oobn int, flags int, addr netip.AddrPort, err error) {
	return x.real().ReadMsgUDPAddrPort(b, oob)
}

func (x *UDPConn) SetReadBuffer(bytes int) error {
	return x.real().SetReadBuffer(bytes)
}

func (x *UDPConn) SetWriteBuffer(bytes int) error {
	return x.real().SetWriteBuffer(bytes)
}

func (x *UDPConn) SyscallConn() (syscall.RawConn, error) {
	return x.real().SyscallConn()
}

func (x *UDPConn) WriteMsgUDP(b []byte, oob []byte, addr *net.UDPAddr) (n int, oobn int, err error) {
	return x.real().WriteMsgUDP(b, oob, addr)
}

func (x *UDPConn) WriteMsgUDPAddrPort(b []byte, oob []byte, addr netip.AddrPort) (n int, oobn int, err error) {
	return x.real().WriteMsgUDPAddrPort(b, oob, addr)
}

func (x *UDPConn) WriteToUDP(b []byte, addr *net.UDPAddr) (int, error) {
	return x.real().WriteToUDP(b, addr)
}

func (x *UDPConn) WriteToUDPAddrPort(b []byte, addr netip.AddrPort) (int, error) {
	return x.real().WriteToUDPAddrPort(b, addr)
}

func (x *UDPConn) real() *net.UDPConn {
	return x.conn.(*net.UDPConn)
}

// ===========================

func newUDPConn(conn *net.UDPConn) *UDPConn {
	value := &UDPConn{}
	initConn(&value.connection, conn)
	return value
}
