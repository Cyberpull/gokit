package cyb

type Client struct {
	Conn
	opts Options
}

func (x *Client) Start() {
	//
}

func NewClient(opts Options) *Client {
	opts.GenerateUUID()

	return &Client{
		opts: opts,
	}
}
