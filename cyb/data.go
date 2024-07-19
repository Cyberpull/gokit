package cyb

import "cyberpull.com/gokit"

type ChannelData struct {
	UUID    string `json:"uuid"`
	Method  string `json:"method"`
	Channel string `json:"channel"`
}

type Data struct {
	Code    int    `json:"code"`
	Content []byte `json:"content"`
}

func (x Data) name() string {
	return "data"
}

func (x Data) prefix() string {
	return "DATA::"
}

func (x Data) GetCode() int {
	return x.Code
}

func (x Data) GetContent() any {
	return x.Content
}

func (x *Data) SetContent(v any) (err error) {
	x.Content, err = toJson(v)
	return
}

func (x Data) Bind(v any) (err error) {
	if x.Content == nil {
		return
	}

	return parseJson(x.Content, v)
}

func (x Data) IsError() bool {
	return x.Code < 200 && x.Code >= 300
}

// ================================

func newData(v any, code ...int) (data *Data) {
	return gokit.PtrOf(mkData(v, code...))
}

func mkData(v any, code ...int) (data Data) {
	var err error

	switch d := v.(type) {
	case Data:
		data = d

	case *Data:
		data = *d

	default:
		if len(code) == 0 {
			code = append(code, 200)
		}

		data.Code = code[0]
		data.Content, err = toJson(gokit.Ptr(v))

		if err != nil {
			data.Content = nil
		}
	}

	switch true {
	case data.Code == 0:
		data.Code = 200

	case data.Code < 200 || data.Code >= 300:
		code[0] = 200
	}

	return
}
