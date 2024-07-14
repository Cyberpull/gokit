package cyb

import (
	"strings"
	"sync"

	"cyberpull.com/gokit"
)

type RouterHander func(ctx *Context) Data

type Router interface {
	Set(method, channel string, handler RouterHander)
}

type ServerRouter struct {
	mutex  sync.Mutex
	mapper map[string]RouterHander
}

func (x *ServerRouter) Set(method, channel string, handler RouterHander) {
	x.mutex.Lock()

	defer x.mutex.Unlock()

	d(x).mapper[x.k(method, channel)] = handler
}

func (x *ServerRouter) Get(method, channel string) (handler RouterHander, ok bool) {
	x.mutex.Lock()

	defer x.mutex.Unlock()

	handler, ok = d(x).mapper[x.k(method, channel)]

	return
}

func (x *ServerRouter) Delete(method, channel string) {
	x.mutex.Lock()

	defer x.mutex.Unlock()

	delete(d(x).mapper, x.k(method, channel))
}

func (x *ServerRouter) k(method, channel string) string {
	method = strings.ToUpper(method)
	return gokit.Join("::", method, channel)
}

func (x *ServerRouter) initialize() {
	if x.mapper == nil {
		x.mapper = make(map[string]RouterHander)
	}
}
