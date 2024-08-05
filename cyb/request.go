package cyb

import (
	"github.com/Cyberpull/gokit"

	"github.com/google/uuid"
)

type Request struct {
	ChannelData
	UUID    string `json:"uuid"`
	Content []byte `json:"content"`
}

func (x Request) name() string {
	return "request"
}

func (x Request) prefix() string {
	return "REQUEST::"
}

func (x *Request) SetContent(v any) (err error) {
	x.Content, err = toJson(v)
	return
}

func (x Request) Bind(v any) (err error) {
	return parseJson(x.Content, v)
}

func (x Request) newResponse(v any, code ...int) *Response {
	return &Response{
		Request:     x,
		UUID:        x.UUID,
		ChannelData: x.ChannelData,
		Data:        mkData(v, code...),
	}
}

// ===============================

func mkRequest(client *Client, data any, cdata ChannelData) (req Request, err error) {
	err = req.SetContent(data)

	if err != nil {
		return
	}

	req.UUID = gokit.Join("::", client.opts.UUID, uuid.NewString())
	req.ChannelData = cdata

	return
}
