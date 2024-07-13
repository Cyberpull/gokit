package cyb

import "context"

type Context struct {
	Request
	context.Context
	in *Inbound
}
