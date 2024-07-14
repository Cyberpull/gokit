package cyb

import "context"

type Context struct {
	context.Context

	req *Request
	in  *Inbound
}

func (x *Context) Update(data Data) (err error) {
	update := &Update{
		Data:        data,
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

func (x *Context) Success(data Data) Data {
	data.Code = 200
	return data
}
