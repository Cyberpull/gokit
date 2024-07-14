package cyb

import (
	"net"
	"sync"

	"cyberpull.com/gokit/graceful"
)

type RouterCallback func(router Router)
type ClientInitCallback func(i *Inbound) (err error)
type InboundPredicate func(i *Inbound) (err error)

type Server struct {
	opts       *Options
	listener   net.Listener
	mutex      sync.Mutex
	mapper     map[string]*Inbound
	clientInit []ClientInitCallback
	router     ServerRouter
}

func (x *Server) Options(opts *Options) {
	x.opts = opts
}

func (x *Server) Router(callbacks ...RouterCallback) {
	for _, callback := range callbacks {
		callback(&x.router)
	}
}

func (x *Server) OnClientInit(callbacks ...ClientInitCallback) {
	x.clientInit = callbacks
}

func (x *Server) Listen() {
	graceful.Run(func(grace graceful.Grace) {
		var err error

		x.listener, err = net.Listen(x.opts.network(), x.opts.address())

		if err != nil {
			return
		}

		for {
			select {
			case <-grace.Done():
				return

			case resp := <-accept(x.listener):
				if err = resp.Error; err != nil {
					return
				}

				inbound := &Inbound{
					Conn:   mkConn(resp.Conn),
					server: x,
				}

				go inbound.Run()
			}
		}
	})
}

func (x *Server) add(i *Inbound) {
	x.mutex.Lock()

	defer x.mutex.Unlock()

	if i.UUID != "" {
		x.mapper[i.UUID] = i
	}
}

func (x *Server) remove(i *Inbound) {
	x.mutex.Lock()

	defer x.mutex.Unlock()

	if i.UUID != "" {
		delete(x.mapper, i.UUID)
	}
}

func (x *Server) each(callback InboundPredicate) {
	x.mutex.Lock()

	defer x.mutex.Unlock()

	for _, i := range x.mapper {
		err := callback(i)

		if err != nil {
			break
		}
	}
}

func (x *Server) initialize() {
	if x.mapper == nil {
		x.mapper = make(map[string]*Inbound)
	}
}
