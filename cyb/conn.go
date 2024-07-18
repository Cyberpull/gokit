package cyb

import (
	"bufio"
	"net"
	"strings"
)

type Conn struct {
	conn   net.Conn
	reader *bufio.Reader
}

func (x *Conn) Write(p []byte) (n int, err error) {
	return x.conn.Write(p)
}

func (x *Conn) WriteLine(p []byte) (n int, err error) {
	return x.Write(append(p, '\n'))
}

func (x *Conn) WriteStringLine(s string) (n int, err error) {
	return x.WriteLine([]byte(s))
}

func (x *Conn) ReadLine() ([]byte, error) {
	return x.reader.ReadBytes('\n')
}

func (x *Conn) ReadString(delim byte) (s string, err error) {
	defer func() {
		s = strings.TrimSuffix(s, string([]byte{delim}))
	}()

	return x.reader.ReadString(delim)
}

func (x *Conn) ReadStringLine() (string, error) {
	return x.ReadString('\n')
}

func (x *Conn) Close() error {
	if x.conn != nil {
		return x.conn.Close()
	}

	return nil
}

// ==================================

func mkConn(conn net.Conn) Conn {
	return Conn{
		conn:   conn,
		reader: bufio.NewReader(conn),
	}
}
