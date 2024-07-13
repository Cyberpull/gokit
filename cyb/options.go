package cyb

import "fmt"

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

func (x Options) address() string {
	switch x.network() {
	case "unix":
		return x.SocketPath

	default:
		return fmt.Sprintf("%v:%v", x.Host, x.Port)
	}
}
