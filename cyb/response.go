package cyb

import (
	"bytes"
	"encoding/json"

	"cyberpull.com/gokit/errors"
)

const responsePrefix string = "RESPONSE::"

type Response struct {
	Data
	ChannelData
	Request Request
}

func (x *Response) ToBytes() (value []byte, err error) {
	value, err = json.Marshal(x)

	if err != nil {
		return
	}

	value = append([]byte(responsePrefix), value...)

	return
}

func (x *Response) ToString() (value string, err error) {
	data, err := x.ToBytes()

	if err != nil {
		return
	}

	value = string(data)

	return
}

func (x *Response) Parse(data []byte) (err error) {
	prefix := []byte(responsePrefix)

	if !bytes.HasPrefix(data, prefix) {
		err = errors.New("Unable to parse response.")
		return
	}

	data = bytes.TrimPrefix(data, prefix)

	err = json.Unmarshal(data, x)

	return
}
