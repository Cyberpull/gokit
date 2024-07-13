package cyb

import (
	"cyberpull.com/gokit/graceful"
)

type Inbound struct {
	Conn
	Info
	server *Server
}

func (x *Inbound) Run() {
	graceful.Run(func(grace graceful.Grace) {
		// Get Client Information
		err := x.getClientInfo()

		if err != nil {
			return
		}

		// Run On Client Init
		err = x.onClientInit()

		if err != nil {
			return
		}

		for {
			select {
			case <-grace.Done():
				return

			case in := <-read(&x.Conn, '\n'):
				if in.Error != nil {
					break
				}

				go x.processRequest(in.Data)
			}
		}
	})
}

func (x *Inbound) getClientInfo() (err error) {
	// TODO: Get Client Info
	return
}

func (x *Inbound) onClientInit() (err error) {
	for _, callback := range x.server.clientInit {
		err = callback(x)

		if err != nil {
			break
		}
	}

	return
}

func (x *Inbound) processRequest(b []byte) (err error) {
	graceful.Run(func(grace graceful.Grace) {
		var req Request

		if err = req.Parse(b); err != nil {
			return
		}

		ctx := &Context{
			in:      x,
			Request: req,
			Context: grace,
		}
	})

	return
}
