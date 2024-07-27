package net

import (
	"context"
	"crypto/tls"
)

type TLSConfig tls.Config

type TLSConn struct {
	connection
}

func (x *TLSConn) ConnectionState() tls.ConnectionState {
	return x.real().ConnectionState()
}

func (x *TLSConn) Handshake() error {
	return x.real().Handshake()
}

func (x *TLSConn) HandshakeContext(ctx context.Context) error {
	return x.real().HandshakeContext(ctx)
}

func (x *TLSConn) NetConn() Conn {
	return newConn(x.real().NetConn())
}

func (x *TLSConn) OCSPResponse() []byte {
	return x.real().OCSPResponse()
}

func (x *TLSConn) VerifyHostname(host string) error {
	return x.real().VerifyHostname(host)
}

func (x *TLSConn) CloseWrite() error {
	return x.real().CloseWrite()
}

func (x *TLSConn) real() *tls.Conn {
	return x.conn.(*tls.Conn)
}

// ===========================

func newTLSConn(conn *tls.Conn) *TLSConn {
	value := &TLSConn{}
	initConn(&value.connection, conn)
	return value
}
