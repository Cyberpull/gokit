package cyb

import (
	"context"
	"encoding/json"
	"net"
	"strings"
	"time"

	"cyberpull.com/gokit"
	"cyberpull.com/gokit/errors"
	"cyberpull.com/gokit/graceful"

	"github.com/google/uuid"
)

type UpdateRouterCallback func(router UpdateRouter)

type Client struct {
	conn           Conn
	srvInfo        Info
	opts           *Options
	router         ClientUpdateRouter
	queue          map[string]chan parsable
	done           chan bool
	canSendRequest bool
	isConnecting   bool
	isConnected    bool
}

func (x *Client) Options(opts *Options) {
	x.opts = opts
}

func (x *Client) Updates(callbacks ...UpdateRouterCallback) {
	for _, callback := range callbacks {
		callback(&x.router)
	}
}

func (x *Client) On(method, channel string, handler UpdateHandler) {
	x.router.Add(method, channel, handler)
}

func (x *Client) Request(method, channel string, data any) (value Data, err error) {
	if !x.canSendRequest {
		err = errors.New("Unable to send request. Please try again later")
		return
	}

	req := Request{
		Content: data,
		ChannelData: ChannelData{
			UUID:    gokit.Join("::", x.opts.UUID, uuid.NewString()),
			Method:  method,
			Channel: channel,
		},
	}

	rawData, err := toBytes(&req)

	if err != nil {
		return
	}

	_, err = x.conn.WriteLine(rawData)

	if err != nil {
		return
	}

	parsableChan := make(chan parsable, 1)

	d(x).queue[req.UUID] = parsableChan

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*30)

	defer func() {
		delete(x.queue, req.UUID)
		cancel()
	}()

	select {
	case <-ctx.Done():
		err = errors.New("Request timed out.")
		return

	case resp := <-parsableChan:
		switch p := resp.(type) {
		case Data:
			value = p

		case Response:
			value = p.Data
		}
	}

	return
}

func (x *Client) Done() chan bool {
	return x.done
}

func (x *Client) Start(autoReconnect ...bool) chan error {
	x.done = make(chan bool, 1)

	defer func() {
		x.done <- true
	}()

	errChan := make(chan error, 1)

	go func() {
		var resultSent bool

		for {
			err := <-x.connect()

			if !resultSent {
				resultSent = true
				errChan <- err
			}

			if err == nil {
				break
			}

			if len(autoReconnect) > 0 && !autoReconnect[0] {
				break
			}

			time.Sleep(time.Second * 10)
		}
	}()

	return errChan
}

func (x *Client) Stop() (err error) {
	x.conn.Close()
	return
}

func (x *Client) connect() (errChan chan error) {
	errChan = make(chan error, 1)

	go func() {
		if x.opts == nil {
			errChan <- errors.New("Invalid options")
			return
		}

		if x.isConnecting {
			errChan <- errors.New("Client already connecting")
			return
		}

		x.isConnecting = true

		defer func() {
			x.isConnecting = false
		}()

		if x.isConnected {
			errChan <- errors.New("Client already connected")
			return
		}

		graceful.Run(func(grace graceful.Grace) {
			conn, err := net.Dial(x.opts.network(), x.opts.address())

			if err != nil {
				errChan <- err
				return
			}

			x.conn = mkConn(conn)
			x.isConnected = true

			defer func() {
				x.isConnected = false
				x.conn = Conn{}
			}()

			if err = x.handshake(); err != nil {
				errChan <- err
				return
			}

			x.canSendRequest = true

			defer func() {
				x.canSendRequest = false
			}()

			errChan <- nil

			for {
				select {
				case <-grace.Done():
					return

				case in := <-gokit.IO.ReadLine(x.conn.conn):
					if in.Error != nil {
						return
					}

					go x.parseInput(in.Data)
				}
			}
		})
	}()

	return
}

func (x *Client) handshake() (err error) {
	msg, err := x.conn.ReadStringLine()

	if err != nil {
		return
	}

	if msg != "CYB::SND" {
		err = errors.New("Invalid HS Received")
		return
	}

	_, err = x.conn.WriteStringLine("CYB::RCV")

	if err != nil {
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

func (x *Client) handshakeProcessName() (err error) {
	msg, err := x.conn.ReadStringLine()

	if err != nil {
		return
	}

	if !strings.HasPrefix(msg, "CYB::NAME=") {
		err = errors.New("Invalid HS Name Received")
		return
	}

	msg = strings.TrimPrefix(msg, "CYB::NAME=")
	err = json.Unmarshal([]byte(msg), &x.srvInfo.Name)

	if err != nil {
		return
	}

	jsonData, err := json.Marshal(x.opts.Name)

	if err != nil {
		return
	}

	_, err = x.conn.WriteStringLine("CYB::NAME=" + string(jsonData))

	return
}

func (x *Client) handshakeProcessDesc() (err error) {
	msg, err := x.conn.ReadStringLine()

	if err != nil {
		return
	}

	if !strings.HasPrefix(msg, "CYB::DESC=") {
		err = errors.New("Invalid HS Description Received")
		return
	}

	msg = strings.TrimPrefix(msg, "CYB::DESC=")
	err = json.Unmarshal([]byte(msg), &x.srvInfo.Description)

	if err != nil {
		return
	}

	jsonData, err := json.Marshal(x.opts.Description)

	if err != nil {
		return
	}

	_, err = x.conn.WriteStringLine("CYB::DESC=" + string(jsonData))

	return
}

func (x *Client) handshakeProcessUUID() (err error) {
	msg, err := x.conn.ReadStringLine()

	if err != nil {
		return
	}

	if !strings.HasPrefix(msg, "CYB::UUID=") {
		err = errors.New("Invalid HS UUID Received")
		return
	}

	msg = strings.TrimPrefix(msg, "CYB::UUID=")
	err = json.Unmarshal([]byte(msg), &x.opts.UUID)

	if err != nil {
		return
	}

	if x.opts.UUID == "" {
		err = errors.New("Empty HS UUID Received")
		return
	}

	_, err = x.conn.WriteStringLine("CYB::UUID::RCV")

	return
}

func (x *Client) parseInput(b []byte) (err error) {
	switch true {
	case hasPrefix(&Error{}, b):
		return x.processError(b)

	case hasPrefix(&Update{}, b):
		return x.processUpdate(b)

	case hasPrefix(&Response{}, b):
		return x.processResponse(b)

	default:
		err = errors.New("Unknown input")
	}

	return
}

func (x *Client) processError(b []byte) (err error) {
	var data Error

	if err = parse(&data, b); err != nil {
		return
	}

	switch true {
	case data.UUID != "":
		if entry, ok := d(x).queue[data.UUID]; ok {
			entry <- data
		}

		delete(x.queue, data.UUID)

	case data.Method != "" && data.Channel != "":
		x.router.Send(data.Method, data.Channel, data.ToData())

	default:
		// Do something
	}

	return
}

func (x *Client) processUpdate(b []byte) (err error) {
	var upd Update

	if err = parse(&upd, b); err != nil {
		return
	}

	x.router.Send(upd.Method, upd.Channel, upd.Data())

	return
}

func (x *Client) processResponse(b []byte) (err error) {
	var resp Response

	if err = parse(&resp, b); err != nil {
		return
	}

	if data, ok := d(x).queue[resp.UUID]; ok {
		data <- resp
	}

	delete(x.queue, resp.UUID)

	return
}

func (x *Client) initialize() {
	if x.queue == nil {
		x.queue = make(map[string]chan parsable)
	}
}
