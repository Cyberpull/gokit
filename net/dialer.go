package net

import (
	"net"
	"time"
)

type Dialer net.Dialer
type IPAddr net.IPAddr
type TCPAddr net.TCPAddr
type UDPAddr net.UDPAddr
type UnixAddr net.UnixAddr

func Dial(network string, address string) (conn Conn, err error) {
	c, err := net.Dial(network, address)

	if err != nil {
		return
	}

	conn = newConn(c)

	return
}

func DialIP(network string, laddr *IPAddr, raddr *IPAddr) (conn *IPConn, err error) {
	c, err := net.DialIP(network, (*net.IPAddr)(laddr), (*net.IPAddr)(raddr))

	if err != nil {
		return
	}

	conn = newIPConn(c)

	return
}

func DialTCP(network string, laddr *TCPAddr, raddr *TCPAddr) (conn *net.TCPConn, err error) {
	c, err := net.DialTCP(network, (*net.TCPAddr)(laddr), (*net.TCPAddr)(raddr))

	if err != nil {
		return
	}

	return
}

func DialTimeout(network string, address string, timeout time.Duration) (conn Conn, err error) {
	c, err := net.DialTimeout(network, address, timeout)

	if err != nil {
		return
	}

	conn = newConn(c)

	return
}

func DialUDP(network string, laddr *net.UDPAddr, raddr *net.UDPAddr) (*net.UDPConn, error) {
	c, err := net.DialUDP()

	if err != nil {
		return
	}

	return
}

func DialUnix(network string, laddr *net.UnixAddr, raddr *net.UnixAddr) (*net.UnixConn, error) {
	c, err := net.DialUnix()

	if err != nil {
		return
	}

	return
}
