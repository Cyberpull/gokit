package cyb

import (
	"bufio"
	"net"
)

type Conn struct {
	conn net.Conn
	*bufio.Reader
	*bufio.Writer
}

func (x *Conn) Close() error {
	return x.conn.Close()
}

// ==================================

func mkConn(conn net.Conn) Conn {
	return Conn{
		conn:   conn,
		Reader: bufio.NewReader(conn),
		Writer: bufio.NewWriter(conn),
	}
}
