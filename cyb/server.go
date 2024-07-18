package cyb

import (
	"net"
	"sync"

	"cyberpull.com/gokit"
	"cyberpull.com/gokit/errors"
	"cyberpull.com/gokit/graceful"
)

type RequestRouterCallback func(router RequestRouter)
type ClientInitCallback func(i *Inbound) (err error)
type InboundPredicate func(i *Inbound) (err error)

type Server struct {
	opts       *Options
	listener   net.Listener
	mutex      sync.Mutex
	mapper     map[string]*Inbound
	clientInit []ClientInitCallback
	router     ServerRequestRouter
	isRunning  bool
	done       chan bool
}

func (x *Server) Options(opts *Options) {
	x.opts = opts
}

func (x *Server) Routes(callbacks ...RequestRouterCallback) {
	for _, callback := range callbacks {
		callback(&x.router)
	}
}

func (x *Server) OnClientInit(callbacks ...ClientInitCallback) {
	x.clientInit = callbacks
}

func (x *Server) Done() chan bool {
	return x.done
}

func (x *Server) Stop() (err error) {
	if x.listener != nil {
		x.listener.Close()
	}

	return
}

func (x *Server) Listen() (errChan chan error) {
	errChan = make(chan error, 1)

	if x.isRunning {
		errChan <- errors.New("Server already running.")
		return
	}

	go graceful.Run(func(grace graceful.Grace) {
		x.isRunning = true
		x.done = make(chan bool, 1)

		defer func() {
			x.isRunning = false
			x.done <- true
		}()

		var err error

		x.opts.freeupAddress()

		x.listener, err = net.Listen(x.opts.network(), x.opts.address())

		if err != nil {
			errChan <- err
			return
		}

		errChan <- nil

		for {
			select {
			case <-grace.Done():
				return

			case resp := <-gokit.Net.Accept(x.listener):
				if err = resp.Error; err != nil {
					return
				}

				inbound := &Inbound{
					Conn:   mkConn(resp.Data),
					server: x,
				}

				go inbound.Run()
			}
		}
	})

	return
}

func (x *Server) add(i *Inbound) {
	x.mutex.Lock()

	defer x.mutex.Unlock()

	if i.UUID != "" {
		d(x).mapper[i.UUID] = i
	}
}

func (x *Server) remove(i *Inbound) {
	x.mutex.Lock()

	defer x.mutex.Unlock()

	if i.UUID != "" {
		delete(d(x).mapper, i.UUID)
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
