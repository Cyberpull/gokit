package cyb

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Cyberpull/gokit"
)

type Options struct {
	Info
	Socket         string
	SocketFileMode os.FileMode
	Host           string
	Port           any
	network        string
	address        string
}

func (x *Options) parse() (err error) {
	x.network, x.address = "", ""

	switch true {
	case x.Socket != "":
		x.network = "unix"
		x.address = gokit.Path.Expand(x.Socket)

		dir := filepath.Dir(x.address)

		if !gokit.Path.IsDir(dir) {
			err = os.MkdirAll(dir, x.perm())

			if err != nil {
				return
			}
		}

	default:
		x.address = fmt.Sprintf("%v:%v", x.Host, x.Port)
		x.network = "tcp"
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
