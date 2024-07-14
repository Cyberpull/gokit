package cyb

import (
	"fmt"
	"log"

	"cyberpull.com/gokit/errors"
	"cyberpull.com/gokit/graceful"
)

type Inbound struct {
	Conn
	Info
	server *Server
}

func (x *Inbound) Run() {
	graceful.Run(func(grace graceful.Grace) {
		// Establish a handshake with client
		err := x.handshake()

		if err != nil {
			return
		}

		x.server.add(x)

		defer func() {
			x.server.remove(x)
		}()

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

func (x *Inbound) handshake() (err error) {
	// TODO: establish handshake with client
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

		if err = parse(req, b); err != nil {
			return
		}

		defer func() {
			rec := recover()

			if rec != nil {
				err = errors.From(rec)
			}

			if err != nil {
				log.Println(err)
			}
		}()

		handler, ok := x.server.router.Get(req.Method, req.Channel)

		if !ok {
			message := fmt.Sprintf(`Action "%v -> %v" not found`, req.Method, req.Channel)
			err = errors.New(message)
			return
		}

		ctx := &Context{
			in:      x,
			req:     &req,
			Context: grace,
		}

		resp := &Response{
			Data:        handler(ctx),
			ChannelData: req.ChannelData,
			Request:     req,
		}

		if resp.Code == 0 {
			resp.Code = 200
		}

		data, err := toBytes(resp)

		if err != nil {
			return
		}

		x.Write(data)
	})

	return
}
