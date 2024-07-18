package gokit

import "net"

type xNet struct {
	//
}

func (x xNet) Accept(listener net.Listener) (value chan IOData[net.Conn]) {
	value = make(chan IOData[net.Conn], 1)

	go func() {
		var resp IOData[net.Conn]
		resp.Data, resp.Error = listener.Accept()
		value <- resp
	}()

	return
}

var Net xNet
