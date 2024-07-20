package cyb

import (
	"context"
)

type Output interface {
	GetCode() int
	GetContent() any
}

type Context struct {
	context.Context

	req   *Request
	in    *Inbound
	queue []*Update
}

func (x *Context) Update(v any, code ...int) (err error) {
	return x.in.Update(x.req.Method, x.req.Channel, v, code...)
}

func (x *Context) UpdateAll(v any, code ...int) (err error) {
	return x.in.UpdateAll(x.req.Method, x.req.Channel, v, code...)
}

func (x *Context) Data(v any, code ...int) Output {
	return newData(v, code...)
}

func (x *Context) Error(v any, code ...int) Output {
	err := newError(v, code...)
	err.ChannelData = x.req.ChannelData
	err.UUID = x.req.UUID
	return err
}

func (x Context) Bind(v any) (err error) {
	return x.req.Bind(v)
}
