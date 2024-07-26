package net

import "net"

type TCPConn struct {
	connection
}

func (x *TCPConn) real() *net.TCPConn {
	return x.connection.conn.(*net.TCPConn)
}

// ===========================

func newTCPConn(conn *net.TCPConn) *TCPConn {
	value := &TCPConn{}
	initConn(&value.connection, conn)
	return value
}

// conn *net.TCPConn
