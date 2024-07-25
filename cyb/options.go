package cyb

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"

	"cyberpull.com/gokit"
	"cyberpull.com/gokit/errors"
)

type Options struct {
	Info
	Socket         string
	SocketFileMode os.FileMode
	network        string
	address        string
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
			x.network = strings.TrimSpace(buff.String())
			buff.Reset()

			if x.network == "" {
				err = errors.New("Invalid socket type")
				return
			}

			x.network = strings.ToLower(x.network)

			continue
		}

		buff.WriteRune(char)
	}

	x.address = strings.TrimSpace(buff.String())

	if x.address == "" {
		err = errors.New("Socket address empty")
		return
	}

	x.address = gokit.Path.Expand(x.address)

	switch x.network {
	case "unix":
		dir := filepath.Dir(x.address)

		_, err = os.Stat(dir)

		if os.IsNotExist(err) {
			err = os.MkdirAll(dir, x.perm())
		}
	}

	return
}

func (x Options) perm() os.FileMode {
	if x.SocketFileMode != 0 {
		return x.SocketFileMode
	}

	return 0775
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
