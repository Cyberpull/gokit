package net

import (
	"bufio"
	"bytes"
	"net"
	"strings"
	"sync"
	"time"
)

type Conn interface {
	// Read reads data from the connection.
	// Read can be made to time out and return an error after a fixed
	// time limit; see SetDeadline and SetReadDeadline.
	Read(p []byte) (n int, err error)
	ReadBytes(delim byte) (b []byte, err error)
	ReadLine() (b []byte, err error)
	ReadString(delim byte) (s string, err error)
	ReadStringLine() (s string, err error)

	// Write writes data to the connection.
	// Write can be made to time out and return an error after a fixed
	// time limit; see SetDeadline and SetWriteDeadline.
	Write(p []byte) (n int, err error)
	WriteLine(p []byte) (n int, err error)
	WriteString(s string) (n int, err error)
	WriteStringLine(s string) (n int, err error)

	// Close closes the connection.
	// Any blocked Read or Write operations will be unblocked and return errors.
	Close() (err error)

	// LocalAddr returns the local network address, if known.
	LocalAddr() net.Addr

	// RemoteAddr returns the remote network address, if known.
	RemoteAddr() net.Addr

	// SetDeadline sets the read and write deadlines associated
	// with the connection. It is equivalent to calling both
	// SetReadDeadline and SetWriteDeadline.
	//
	// A deadline is an absolute time after which I/O operations
	// fail instead of blocking. The deadline applies to all future
	// and pending I/O, not just the immediately following call to
	// Read or Write. After a deadline has been exceeded, the
	// connection can be refreshed by setting a deadline in the future.
	//
	// If the deadline is exceeded a call to Read or Write or to other
	// I/O methods will return an error that wraps os.ErrDeadlineExceeded.
	// This can be tested using errors.Is(err, os.ErrDeadlineExceeded).
	// The error's Timeout method will return true, but note that there
	// are other possible errors for which the Timeout method will
	// return true even if the deadline has not been exceeded.
	//
	// An idle timeout can be implemented by repeatedly extending
	// the deadline after successful Read or Write calls.
	//
	// A zero value for t means I/O operations will not time out.
	SetDeadline(t time.Time) error

	// SetReadDeadline sets the deadline for future Read calls
	// and any currently-blocked Read call.
	// A zero value for t means Read will not time out.
	SetReadDeadline(t time.Time) error

	// SetWriteDeadline sets the deadline for future Write calls
	// and any currently-blocked Write call.
	// Even if write times out, it may return n > 0, indicating that
	// some of the data was successfully written.
	// A zero value for t means Write will not time out.
	SetWriteDeadline(t time.Time) error
}

type connection struct {
	wm     sync.Mutex
	rm     sync.Mutex
	conn   net.Conn
	reader *bufio.Reader
	trim   bool
}

func (x *connection) Read(p []byte) (n int, err error) {
	x.rm.Lock()

	defer x.rm.Unlock()

	return x.conn.Read(p)
}

func (x *connection) ReadBytes(delim byte) (b []byte, err error) {
	defer func() {
		if x.trim {
			b = bytes.TrimSuffix(b, []byte{delim})
		}
	}()

	b, err = x.reader.ReadBytes(delim)

	return
}

func (x *connection) ReadLine() (b []byte, err error) {
	return x.ReadBytes('\n')
}

func (x *connection) ReadString(delim byte) (s string, err error) {
	defer func() {
		if x.trim {
			s = strings.TrimSuffix(s, string([]byte{delim}))
		}
	}()

	s, err = x.reader.ReadString(delim)

	return
}

func (x *connection) ReadStringLine() (s string, err error) {
	return x.ReadString('\n')
}

func (x *connection) Write(p []byte) (n int, err error) {
	x.wm.Lock()

	defer x.wm.Unlock()

	return x.conn.Write(p)
}

func (x *connection) WriteLine(p []byte) (n int, err error) {
	return x.Write(append(p, '\n'))
}

func (x *connection) WriteString(s string) (n int, err error) {
	return x.Write([]byte(s))
}

func (x *connection) WriteStringLine(s string) (n int, err error) {
	return x.WriteLine([]byte(s))
}

func (x *connection) Close() (err error) {
	if x.conn != nil {
		err = x.conn.Close()
	}

	return
}

func (x *connection) LocalAddr() net.Addr {
	return x.conn.LocalAddr()
}

func (x *connection) RemoteAddr() net.Addr {
	return x.conn.RemoteAddr()
}

func (x *connection) SetDeadline(t time.Time) error {
	return x.conn.SetDeadline(t)
}

func (x *connection) SetReadDeadline(t time.Time) error {
	return x.conn.SetReadDeadline(t)
}

func (x *connection) SetWriteDeadline(t time.Time) error {
	return x.conn.SetWriteDeadline(t)
}

func (x *connection) packet() net.PacketConn {
	return x.conn.(net.PacketConn)
}

// ==================================

func newConn(conn net.Conn) Conn {
	value := &connection{}
	initConn(value, conn, true)
	return value
}

func initConn(x *connection, conn net.Conn, trim ...bool) {
	x.conn = conn
	x.reader = bufio.NewReader(x)
	x.trim = len(trim) > 0 && trim[0]
}
