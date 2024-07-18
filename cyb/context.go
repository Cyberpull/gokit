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

	req *Request
	in  *Inbound
}

func (x *Context) Update(v any, code ...int) (err error) {
	data := mkData(v, code...)

	update := &Update{
		Code:        data.Code,
		Content:     data.Content,
		ChannelData: x.req.ChannelData,
	}

	value, err := toBytes(update)

	if err != nil {
		return
	}

	x.in.server.each(func(i *Inbound) (err error) {
		i.Write(value)
		return
	})

	return
}

func (x *Context) Data(v any, code ...int) Output {
	return newData(v, code...)
}

func (x *Context) Error(v any, code ...int) Output {
	err := newError(v, code...)
	err.ChannelData = x.req.ChannelData
	return err
}
