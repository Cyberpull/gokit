package cyb

type Request struct {
	Data
	ChannelData
}

func (x Request) name() string {
	return "request"
}

func (x Request) prefix() string {
	return "REQUEST::"
}
