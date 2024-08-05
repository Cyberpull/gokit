package net

import (
	"crypto/tls"
	"net"

	"github.com/Cyberpull/gokit"
)

type Listener interface {
	// Accept waits for and returns the next connection to the listener.
	Accept() (conn Conn, err error)

	// Accept waits for and returns the next connection to the listener as an IOData channel.
	AcceptChan() (value chan gokit.IOData[Conn])

	// Close closes the listener.
	// Any blocked Accept operations will be unblocked and return errors.
	Close() error

	// Addr returns the listener's network address.
	Addr() Addr
}

type listener struct {
	listener net.Listener
}

func (x *listener) Accept() (conn Conn, err error) {
	c, err := x.listener.Accept()

	if err != nil {
		return
	}

	conn = newConn(c)

	return
}

func (x *listener) AcceptChan() (value chan gokit.IOData[Conn]) {
	value = make(chan gokit.IOData[Conn], 1)

	go func() {
		var resp gokit.IOData[Conn]
		resp.Data, resp.Error = x.Accept()
		value <- resp
	}()

	return
}

func (x *listener) Close() error {
	return x.listener.Close()
}

func (x *listener) Addr() Addr {
	return x.listener.Addr()
}

// ======================================

func Listen(network string, address string) (value Listener, err error) {
	l, err := net.Listen(network, address)

	if err != nil {
		return
	}

	value = &listener{listener: l}

	return
}

func ListenTLS(network string, laddr string, config *TLSConfig) (value Listener, err error) {
	l, err := tls.Listen(network, laddr, (*tls.Config)(config))

	if err != nil {
		return
	}

	value = &listener{listener: l}

	return
}
