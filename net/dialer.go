package net

import (
	"crypto/tls"
	"net"
	"time"
)

type Dialer net.Dialer
type TLSDialer tls.Dialer

type Addr net.Addr
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

func DialTCP(network string, laddr *TCPAddr, raddr *TCPAddr) (conn *TCPConn, err error) {
	c, err := net.DialTCP(network, (*net.TCPAddr)(laddr), (*net.TCPAddr)(raddr))

	if err != nil {
		return
	}

	conn = newTCPConn(c)

	return
}

func DialTLS(network string, addr string, config *TLSConfig) (conn *TLSConn, err error) {
	c, err := tls.Dial(network, addr, (*tls.Config)(config))

	if err != nil {
		return
	}

	conn = newTLSConn(c)

	return
}

func DialTLSWithDialer(dialer *Dialer, network string, addr string, config *TLSConfig) (conn *TLSConn, err error) {
	c, err := tls.DialWithDialer((*net.Dialer)(dialer), network, addr, (*tls.Config)(config))

	if err != nil {
		return
	}

	conn = newTLSConn(c)

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

func DialUDP(network string, laddr *UDPAddr, raddr *UDPAddr) (conn *UDPConn, err error) {
	c, err := net.DialUDP(network, (*net.UDPAddr)(laddr), (*net.UDPAddr)(raddr))

	if err != nil {
		return
	}

	conn = newUDPConn(c)

	return
}

func DialUnix(network string, laddr *UnixAddr, raddr *UnixAddr) (conn *UnixConn, err error) {
	c, err := net.DialUnix(network, (*net.UnixAddr)(laddr), (*net.UnixAddr)(raddr))

	if err != nil {
		return
	}

	conn = newUnixConn(c)

	return
}
