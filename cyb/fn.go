package cyb

import (
	"net"
)

type ChanData struct {
	Data  []byte
	Error error
}

type ChanListenData struct {
	Conn  net.Conn
	Error error
}

func read(conn *Conn, delim byte) (resp chan ChanData) {
	resp = make(chan ChanData, 1)

	go func() {
		var x ChanData

		x.Data, x.Error = conn.ReadBytes(delim)

		resp <- x
	}()

	return
}

func accept(listener net.Listener) (data chan ChanListenData) {
	data = make(chan ChanListenData, 1)

	go func() {
		var resp ChanListenData

		resp.Conn, resp.Error = listener.Accept()

		data <- resp
	}()

	return
}
