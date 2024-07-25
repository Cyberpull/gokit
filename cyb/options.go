package cyb

import (
	"bytes"
	"os"

	"cyberpull.com/gokit/errors"
)

type Options struct {
	Info
	Socket  string
	network string
	address string
}

func (x *Options) parse() (err error) {
	if x.Socket == "" {
		err = errors.New("Socket empty")
		return
	}

	x.network, x.address = "", ""

	var buff bytes.Buffer

	for _, char := range x.Socket {
		if char == ':' && x.network == "" {
			x.network = buff.String()
			buff.Reset()

			if x.network == "" {
				err = errors.New("Invalid socket type")
				return
			}

			continue
		}

		buff.WriteRune(char)
	}

	if buff.Len() == 0 {
		err = errors.New("Socket address empty")
		return
	}

	x.address = buff.String()

	return
}

func (x Options) freeupAddress() {
	switch x.network {
	case "unix":
		info, err := os.Stat(x.address)

		if err == nil && !info.IsDir() {
			os.Remove(x.address)
		}

	default:
		// Free up port
	}
}
