package cyb

type Update struct {
	ChannelData
	Code    int    `json:"code"`
	Content []byte `json:"content"`
}

func (x Update) name() string {
	return "update"
}

func (x Update) prefix() string {
	return "UPDATE::"
}

func (x *Update) SetContent(v any) (err error) {
	x.Content, err = toJson(v)
	return
}

func (x Update) Data() Data {
	return Data{
		Code:    x.Code,
		Content: x.Content,
	}
}
