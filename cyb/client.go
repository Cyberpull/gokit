package cyb

import (
	"context"
	"strings"
	"sync"
	"time"

	"github.com/Cyberpull/gokit"
	"github.com/Cyberpull/gokit/errors"
	"github.com/Cyberpull/gokit/graceful"
	"github.com/Cyberpull/gokit/net"
)

type UpdateRouterCallback func(router UpdateRouter)

type Client struct {
	conn           net.Conn
	srvInfo        Info
	opts           *Options
	mutex          sync.Mutex
	router         ClientUpdateRouter
	queue          map[string]chan parsable
	bootCallbacks  []BootCallback
	authCallbacks  []AuthCallback
	requestTimeout int
	canSendRequest bool
	isConnecting   bool
}

func (x *Client) Boot(callbacks ...BootCallback) {
	x.bootCallbacks = callbacks
}

func (x *Client) Auth(callbacks ...AuthCallback) {
	x.authCallbacks = callbacks
}

// Set Timeout for requests (in seconds).
// Defaults to 30 seconds
func (x *Client) SetRequestTimeout(timeout int) {
	x.requestTimeout = timeout
}

func (x *Client) Updates(callbacks ...UpdateRouterCallback) {
	for _, callback := range callbacks {
		callback(&x.router)
	}
}

func (x *Client) On(method, channel string, handler UpdateHandler) {
	x.router.On(method, channel, handler)
}

func (x *Client) Request(method, channel string, data any) (value Data, err error) {
	if !x.canSendRequest {
		err = errors.New("Unable to send request. Please try again later")
		return
	}

	req, err := mkRequest(x, data, ChannelData{
		Method:  method,
		Channel: channel,
	})

	if err != nil {
		return
	}

	rawData, err := toBytes(&req)

	if err != nil {
		return
	}

	_, err = x.conn.WriteLine(rawData)

	if err != nil {
		return
	}

	resp := <-x.getResponse(&req)

	if err = resp.Error; err != nil {
		return
	}

	value = resp.Data

	return
}

func (x *Client) sendResponse(reqId string, resp parsable) {
	x.mutex.Lock()

	defer x.mutex.Unlock()

	if respChan, ok := d(x).queue[reqId]; ok {
		respChan <- resp
	}

	delete(x.queue, reqId)
}

func (x *Client) addResponseChannel(reqId string, respChan chan parsable) {
	x.mutex.Lock()

	defer x.mutex.Unlock()

	d(x).queue[reqId] = respChan
}

func (x *Client) getResponse(req *Request) (value chan gokit.IOData[Data]) {
	value = make(chan gokit.IOData[Data], 1)

	go func() {
		var resp gokit.IOData[Data]

		respChan := make(chan parsable, 1)
		x.addResponseChannel(req.UUID, respChan)

		timeout := time.Duration(d(x).requestTimeout)
		ctx, cancel := context.WithTimeout(context.TODO(), time.Second*timeout)

		defer func() {
			value <- resp
			cancel()
		}()

		select {
		case <-ctx.Done():
			resp.Error = errors.New("Request timed out.")
			return

		case data := <-respChan:
			switch d := data.(type) {
			case Error:
				resp.Error = d.Error()

			case Data:
				resp.Data = d

			case Response:
				resp.Data = d.Data
			}
		}
	}()

	return
}

func (x *Client) Start(opts *Options) chan error {
	errChan := make(chan error, 1)

	go graceful.Run(func(grace graceful.Grace) {
		var resultSent bool

		for {
			select {
			case <-grace.Done():
				return

			case err := <-x.Connect(opts):
				if !resultSent {
					resultSent = true
					errChan <- err
				}

				if err == nil {
					x.Run()
				}

				time.Sleep(time.Second * 10)
			}
		}
	})

	return errChan
}

func (x *Client) isConnected() bool {
	return x.conn != nil
}

func (x *Client) Stop() (err error) {
	if x.conn != nil {
		x.conn.Close()
		x.conn = nil
	}

	x.opts = nil
	x.srvInfo = Info{}

	x.canSendRequest = false
	x.isConnecting = false

	return
}

func (x *Client) Connect(opts *Options) (errChan chan error) {
	errChan = make(chan error, 1)

	graceful.Run(func(grace graceful.Grace) {
		var err error

		defer func() {
			errChan <- err

			if err != nil {
				x.Stop()
			}
		}()

		if opts == nil {
			err = errors.New("Invalid options")
			return
		}

		if err = opts.parse(); err != nil {
			return
		}

		x.opts = opts

		if x.isConnecting {
			err = errors.New("Client already connecting")
			return
		}

		x.isConnecting = true

		defer func() {
			x.isConnecting = false
		}()

		if x.isConnected() {
			err = errors.New("Client already connected")
			return
		}

		err = x.execBoot()

		if err != nil {
			return
		}

		x.conn, err = net.Dial(x.opts.network, x.opts.address)

		if err != nil {
			return
		}

		err = x.handshake()

		if err != nil {
			return
		}

		err = x.execAuth()
	})

	return
}

func (x *Client) Run() (err error) {
	graceful.Run(func(grace graceful.Grace) {
		defer x.Stop()

		if !x.isConnected() {
			err = errors.New("Client not connected")
			return
		}

		x.canSendRequest = true

		for {
			select {
			case <-grace.Done():
				return

			case in := <-gokit.IO.ReadLine(x.conn):
				if in.Error != nil {
					return
				}

				go x.parseInput(in.Data)
			}
		}
	})

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
	err = parseJson([]byte(msg), &x.srvInfo.Name)

	if err != nil {
		return
	}

	jsonData, err := toJson(x.opts.Name)

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
	err = parseJson([]byte(msg), &x.srvInfo.Description)

	if err != nil {
		return
	}

	jsonData, err := toJson(x.opts.Description)

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
	err = parseJson([]byte(msg), &x.opts.UUID)

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
		x.sendResponse(data.UUID, data)

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

	// x.conn.WriteStringLine("OK::" + upd.UUID)

	x.router.Send(upd.Method, upd.Channel, upd.Data())

	return
}

func (x *Client) processResponse(b []byte) (err error) {
	var resp Response

	if err = parse(&resp, b); err != nil {
		return
	}

	x.sendResponse(resp.UUID, resp)

	return
}

func (x *Client) execBoot() (err error) {
	for _, callback := range x.bootCallbacks {
		err = callback()

		if err != nil {
			break
		}
	}

	return
}

func (x *Client) execAuth() (err error) {
	for _, callback := range x.authCallbacks {
		err = callback(x.conn)

		if err != nil {
			break
		}
	}

	return
}

func (x *Client) initialize() {
	if x.queue == nil {
		x.queue = make(map[string]chan parsable)
	}

	if x.requestTimeout == 0 {
		x.requestTimeout = 30
	}
}
