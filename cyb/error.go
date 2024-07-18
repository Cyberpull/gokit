package cyb

import (
	"cyberpull.com/gokit"
	"cyberpull.com/gokit/errors"
)

type Error struct {
	ChannelData
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (x Error) GetCode() int {
	return x.Code
}

func (x Error) GetContent() any {
	return x.Message
}

func (x Error) name() string {
	return "error"
}

func (x Error) prefix() string {
	return "ERROR::"
}

func (x Error) ToData() Data {
	return Data{Code: x.Code, Content: x.Message}
}

// ================================

func newError(v any, code ...int) (data *Error) {
	return gokit.PtrOf(mkError(v, code...))
}

func mkError(v any, code ...int) (data Error) {
	switch true {
	case len(code) == 0:
		code = append(code, 500)

	case code[0] == 0, code[0] >= 200 && code[0] < 300:
		code[0] = 500
	}

	switch d := v.(type) {
	case *errors.Error:
		data.Code = d.Code()
		data.Message = d.Error()

	case error:
		data.Code = code[0]
		data.Message = d.Error()

	case string:
		data.Code = code[0]
		data.Message = d

	case []byte:
		data.Code = code[0]
		data.Message = string(d)

	default:
		data.Code = code[0]
		data.Message = "An unknown error occurred"
	}

	return
}
