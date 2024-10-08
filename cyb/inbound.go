package cyb

import (
	"fmt"
	"log"
	"strings"

	"github.com/Cyberpull/gokit"
	"github.com/Cyberpull/gokit/errors"
	"github.com/Cyberpull/gokit/graceful"
	"github.com/Cyberpull/gokit/net"

	"github.com/google/uuid"
)

type InboundConnection interface {
	Update(method, channel string, v any, code ...int) (err error)
}

type Inbound struct {
	Info
	conn     net.Conn
	server   *Server
	updQueue map[string]chan string
}

func (x *Inbound) Run() {
	graceful.Run(func(grace graceful.Grace) {
		// Establish a handshake with client
		err := x.handshake()

		if err != nil {
			return
		}

		err = x.execAuth()

		if err != nil {
			return
		}

		x.server.add(x)

		defer func() {
			x.server.remove(x)
		}()

		// Run On Client Init
		err = x.execClientInit()

		if err != nil {
			return
		}

		for {
			select {
			case <-grace.Done():
				return

			case in := <-gokit.IO.ReadLine(x.conn):
				if in.Error != nil {
					return
				}

				go x.processRequest(in.Data)
			}
		}
	})
}

func (x *Inbound) Update(method, channel string, v any, code ...int) (err error) {
	data := mkData(v, code...)

	update := &Update{
		Code:    data.Code,
		Content: data.Content,
		ChannelData: ChannelData{
			Method:  method,
			Channel: channel,
		},
	}

	value, err := toBytes(update)

	if err != nil {
		return
	}

	_, err = x.conn.WriteLine(value)

	return
}

func (x *Inbound) UpdateAll(method, channel string, v any, code ...int) (err error) {
	data := mkData(v, code...)

	update := &Update{
		Code:    data.Code,
		Content: data.Content,
		ChannelData: ChannelData{
			Method:  method,
			Channel: channel,
		},
	}

	value, err := toBytes(update)

	if err != nil {
		return
	}

	go x.server.forEach(func(i *Inbound) (err error) {
		i.conn.WriteLine(value)
		return
	})

	// go func() {
	// 	for _, in := range x.server.mapper {
	// 		in.conn.WriteLine(value)
	// 	}
	// }()

	return
}

func (x *Inbound) handshake() (err error) {
	_, err = x.conn.WriteStringLine("CYB::SND")

	if err != nil {
		return
	}

	resp, err := x.conn.ReadStringLine()

	if err != nil {
		return
	}

	if resp != "CYB::RCV" {
		err = errors.New("Invalid HS Response")
		return
	}

	err = x.handshakeProcessName()

	if err != nil {
		return
	}

	err = x.handshakeProcessDesc()

	if err != nil {
		return
	}

	err = x.handshakeProcessUUID()

	return
}

func (x *Inbound) handshakeProcessName() (err error) {
	jsonData, err := toJson(x.server.opts.Name)

	if err != nil {
		return
	}

	_, err = x.conn.WriteStringLine("CYB::NAME=" + string(jsonData))

	if err != nil {
		return
	}

	resp, err := x.conn.ReadStringLine()

	if err != nil {
		return
	}

	if !strings.HasPrefix(resp, "CYB::NAME=") {
		err = errors.New("Invalid HS Name Received")
		return
	}

	resp = strings.TrimPrefix(resp, "CYB::NAME=")

	var name string

	if err = parseJson([]byte(resp), &name); err != nil {
		return
	}

	x.Name = name

	return
}

func (x *Inbound) handshakeProcessDesc() (err error) {
	jsonData, err := toJson(x.server.opts.Description)

	if err != nil {
		return
	}

	_, err = x.conn.WriteStringLine("CYB::DESC=" + string(jsonData))

	if err != nil {
		return
	}

	resp, err := x.conn.ReadStringLine()

	if err != nil {
		return
	}

	if !strings.HasPrefix(resp, "CYB::DESC=") {
		err = errors.New("Invalid HS Description Received")
		return
	}

	resp = strings.TrimPrefix(resp, "CYB::DESC=")

	var desc string

	if err = parseJson([]byte(resp), &desc); err != nil {
		return
	}

	x.Description = desc

	return
}

func (x *Inbound) handshakeProcessUUID() (err error) {
	clientUUID := uuid.NewString()
	jsonData, err := toJson(clientUUID)

	if err != nil {
		return
	}

	_, err = x.conn.WriteStringLine("CYB::UUID=" + string(jsonData))

	if err != nil {
		return
	}

	resp, err := x.conn.ReadStringLine()

	if err != nil {
		return
	}

	if resp != "CYB::UUID::RCV" {
		err = errors.New("Invalid HS UUID Response Received")
		return
	}

	x.UUID = clientUUID

	return
}

func (x *Inbound) processRequest(b []byte) (err error) {
	graceful.Run(func(grace graceful.Grace) {
		var req Request

		if err = parse(&req, b); err != nil {
			return
		}

		defer func() {
			rec := recover()

			if rec != nil {
				err = errors.From(rec)
			}

			if err != nil {
				resp := mkError(err)

				if data, e := toBytes(resp); e == nil {
					x.conn.WriteLine(data)
				}

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
			queue:   make([]*Update, 0),
			Context: grace,
		}

		result := handler(ctx)

		switch d := result.(type) {
		case *Error:
			data, err := toBytes(d)

			if err != nil {
				return
			}

			x.conn.WriteLine(data)

		case *Data:
			resp := req.newResponse(d)

			data, err := toBytes(resp)

			if err != nil {
				return
			}

			x.conn.WriteLine(data)

		default:
			err = errors.New("Unknown response")
		}
	})

	return
}

func (x *Inbound) execAuth() (err error) {
	for _, callback := range x.server.authCallbacks {
		err = callback(x.conn)

		if err != nil {
			break
		}
	}

	return
}

func (x *Inbound) execClientInit() (err error) {
	for _, callback := range x.server.clientInitCallbacks {
		err = callback(x)

		if err != nil {
			break
		}
	}

	return
}
