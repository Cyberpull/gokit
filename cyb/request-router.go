package cyb

import (
	"strings"
	"sync"

	"cyberpull.com/gokit"
)

type RequestHander func(ctx *Context) Output

type RequestRouter interface {
	Set(method, channel string, handler RequestHander)
}

type ServerRequestRouter struct {
	mutex  sync.Mutex
	mapper map[string]RequestHander
}

func (x *ServerRequestRouter) Set(method, channel string, handler RequestHander) {
	x.mutex.Lock()

	defer x.mutex.Unlock()

	d(x).mapper[x.k(method, channel)] = handler
}

func (x *ServerRequestRouter) Get(method, channel string) (handler RequestHander, ok bool) {
	x.mutex.Lock()

	defer x.mutex.Unlock()

	handler, ok = d(x).mapper[x.k(method, channel)]

	return
}

func (x *ServerRequestRouter) Delete(method, channel string) {
	x.mutex.Lock()

	defer x.mutex.Unlock()

	delete(d(x).mapper, x.k(method, channel))
}

func (x *ServerRequestRouter) k(method, channel string) string {
	method = strings.ToUpper(method)
	return gokit.Join("::", method, channel)
}

func (x *ServerRequestRouter) initialize() {
	if x.mapper == nil {
		x.mapper = make(map[string]RequestHander)
	}
}
