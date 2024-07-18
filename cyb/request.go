package cyb

type Request struct {
	ChannelData
	Content any `json:"content"`
}

func (x Request) name() string {
	return "request"
}

func (x Request) prefix() string {
	return "REQUEST::"
}

func (x Request) newResponse(v any, code ...int) *Response {
	return &Response{
		Request:     x,
		ChannelData: x.ChannelData,
		Data:        mkData(v, code...),
	}
}
