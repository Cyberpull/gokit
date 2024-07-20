package cyb

type Response struct {
	Data
	ChannelData
	UUID    string `json:"uuid"`
	Request Request
}

func (x Response) name() string {
	return "response"
}

func (x Response) prefix() string {
	return "RESPONSE::"
}
