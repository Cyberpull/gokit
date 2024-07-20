package cyb

import (
	"log"
	"strings"
	"sync"

	"cyberpull.com/gokit"
)

type UpdateHandler func(data OutputData)

type UpdateRouter interface {
	On(method, channel string, handler UpdateHandler)
}

type ClientUpdateRouter struct {
	mutex  sync.Mutex
	mapper map[string][]UpdateHandler
}

func (x *ClientUpdateRouter) On(method, channel string, handler UpdateHandler) {
	x.mutex.Lock()

	defer x.mutex.Unlock()

	key := x.k(method, channel)

	if _, ok := d(x).mapper[key]; !ok {
		x.mapper[key] = make([]UpdateHandler, 0)
	}

	x.mapper[key] = append(x.mapper[key], handler)
}

func (x *ClientUpdateRouter) Send(method, channel string, data Data) {
	x.mutex.Lock()

	defer x.mutex.Unlock()

	key := x.k(method, channel)

	if _, ok := d(x).mapper[key]; !ok {
		return
	}

	defer func() {
		rec := recover()

		if rec != nil {
			log.Println(rec)
		}
	}()

	for _, handler := range x.mapper[key] {
		handler(data)
	}
}

func (x *ClientUpdateRouter) Clear(method, channel string) {
	x.mutex.Lock()

	defer x.mutex.Unlock()

	delete(d(x).mapper, x.k(method, channel))
}

func (x *ClientUpdateRouter) k(method, channel string) string {
	method = strings.ToUpper(method)
	return gokit.Join("::", method, channel)
}

func (x *ClientUpdateRouter) initialize() {
	if x.mapper == nil {
		x.mapper = make(map[string][]UpdateHandler)
	}
}
