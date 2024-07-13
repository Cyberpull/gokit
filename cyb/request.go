package cyb

import (
	"bytes"
	"encoding/json"

	"cyberpull.com/gokit/errors"
)

const requestPrefix string = "REQUEST::"

type Request struct {
	Data
	ChannelData
}

func (x *Request) ToBytes() (value []byte, err error) {
	value, err = json.Marshal(x)

	if err != nil {
		return
	}

	value = append([]byte(requestPrefix), value...)

	return
}

func (x *Request) ToString() (value string, err error) {
	data, err := x.ToBytes()

	if err != nil {
		return
	}

	value = string(data)

	return
}

func (x *Request) Parse(data []byte) (err error) {
	prefix := []byte(requestPrefix)

	if !bytes.HasPrefix(data, prefix) {
		err = errors.New("Unable to parse request.")
		return
	}

	data = bytes.TrimPrefix(data, prefix)

	err = json.Unmarshal(data, x)

	return
}
