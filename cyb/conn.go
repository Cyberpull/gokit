package cyb

import (
	"bufio"
	"bytes"
	"net"
	"strings"
	"sync"
)

type Connection interface {
	ReadBytes(delim byte) (b []byte, err error)
	ReadLine() (b []byte, err error)
	ReadString(delim byte) (s string, err error)
	ReadStringLine() (s string, err error)

	Write(p []byte) (n int, err error)
	WriteLine(p []byte) (n int, err error)
	WriteString(s string) (n int, err error)
	WriteStringLine(s string) (n int, err error)
}

type Conn struct {
	wm     sync.Mutex
	rm     sync.Mutex
	conn   net.Conn
	reader *bufio.Reader
}

func (x *Conn) Read(p []byte) (n int, err error) {
	x.rm.Lock()

	defer x.rm.Unlock()

	return x.conn.Read(p)
}

func (x *Conn) ReadBytes(delim byte) (b []byte, err error) {
	defer func() {
		b = bytes.TrimSuffix(b, []byte{delim})
	}()

	b, err = x.reader.ReadBytes(delim)

	return
}

func (x *Conn) ReadLine() (b []byte, err error) {
	return x.ReadBytes('\n')
}

func (x *Conn) ReadString(delim byte) (s string, err error) {
	defer func() {
		s = strings.TrimSuffix(s, string([]byte{delim}))
	}()

	s, err = x.reader.ReadString(delim)

	return
}

func (x *Conn) ReadStringLine() (s string, err error) {
	return x.ReadString('\n')
}

func (x *Conn) Write(p []byte) (n int, err error) {
	x.wm.Lock()

	defer x.wm.Unlock()

	return x.conn.Write(p)
}

func (x *Conn) WriteLine(p []byte) (n int, err error) {
	return x.Write(append(p, '\n'))
}

func (x *Conn) WriteString(s string) (n int, err error) {
	return x.Write([]byte(s))
}

func (x *Conn) WriteStringLine(s string) (n int, err error) {
	return x.WriteLine([]byte(s))
}

func (x *Conn) Close() error {
	if x.conn != nil {
		return x.conn.Close()
	}

	return nil
}

// ==================================

func newConn(conn net.Conn) (value *Conn) {
	value = &Conn{conn: conn}
	value.reader = bufio.NewReader(value)
	return
}
