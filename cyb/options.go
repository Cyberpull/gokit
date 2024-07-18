package cyb

import (
	"fmt"
	"os"
)

type Options struct {
	Info
	Host       string
	Port       any
	Network    string
	SocketPath string
}

func (x Options) network() string {
	if x.Network != "" {
		return x.Network
	}

	return "tcp"
}

func (x Options) freeupAddress() {
	switch x.network() {
	case "unix":
		info, err := os.Stat(x.SocketPath)

		if err == nil && !info.IsDir() {
			os.Remove(x.SocketPath)
		}

	default:
		// Free up port
	}
}

func (x Options) address() string {
	switch x.network() {
	case "unix":
		return x.SocketPath

	default:
		return fmt.Sprintf("%v:%v", x.Host, x.Port)
	}
}
