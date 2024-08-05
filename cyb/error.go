package cyb

import (
	"github.com/Cyberpull/gokit"
	"github.com/Cyberpull/gokit/errors"
)

type Error struct {
	ChannelData
	UUID    string `json:"uuid"`
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

func (x Error) ToData() (data Data) {
	var err error

	data.Code = x.Code

	if data.Content, err = toJson(&x); err != nil {
		data.Content = nil
	}

	return
}

func (x Error) Error() (err error) {
	return errors.New(x.Message, x.Code)
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
