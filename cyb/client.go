package cyb

import "sync"

type Client struct {
	Conn
	mutex sync.Mutex
	opts  Options
}

func (x *Client) Options(opts Options) {
	x.opts = opts
}

func (x *Client) Update() {
	//
}

func (x *Client) Start() {
	x.opts.GenerateUUID()
}

func (x *Client) handshake() (err error) {
	// TODO: establish handshake with server
	return
}

func (x *Client) initialize() {
	//
}
