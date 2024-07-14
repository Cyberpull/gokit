package cyb

type Client struct {
	Conn
	opts Options
}

func (x *Client) Start() {
	//
}

func (x *Client) handshake() (err error) {
	// TODO: establish handshake with server
	return
}

func (x *Client) initialize() {
	//
}

func NewClient(opts Options) *Client {
	opts.GenerateUUID()

	return &Client{
		opts: opts,
	}
}
