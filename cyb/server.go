package cyb

import (
	"sync"

	"github.com/Cyberpull/gokit/errors"
	"github.com/Cyberpull/gokit/graceful"
	"github.com/Cyberpull/gokit/net"
)

type BootCallback func() (err error)
type AuthCallback func(conn Connection) (err error)
type RequestRouterCallback func(router RequestRouter)
type ClientInitCallback func(i InboundConnection) (err error)
type InboundPredicate func(i *Inbound) (err error)

type Server struct {
	opts                *Options
	listener            net.Listener
	mutex               sync.Mutex
	mapper              map[string]*Inbound
	clientInitCallbacks []ClientInitCallback
	bootCallbacks       []BootCallback
	authCallbacks       []AuthCallback
	router              ServerRequestRouter
}

func (x *Server) Boot(callbacks ...BootCallback) {
	x.bootCallbacks = callbacks
}

func (x *Server) Auth(callbacks ...AuthCallback) {
	x.authCallbacks = callbacks
}

func (x *Server) Routes(callbacks ...RequestRouterCallback) {
	for _, callback := range callbacks {
		callback(&x.router)
	}
}

func (x *Server) OnClientInit(callbacks ...ClientInitCallback) {
	x.clientInitCallbacks = callbacks
}

func (x *Server) Stop() (err error) {
	if x.listener != nil {
		x.listener.Close()
		x.listener = nil
	}

	return
}

func (x *Server) isListening() bool {
	return x.listener != nil
}

func (x *Server) Listen(opts *Options) (err error) {
	err = <-x.Connect(opts)

	if err != nil {
		return
	}

	err = x.Run()

	return
}

func (x *Server) Connect(opts *Options) (errChan chan error) {
	errChan = make(chan error, 1)

	go graceful.Run(func(grace graceful.Grace) {
		var err error

		defer func() {
			errChan <- err

			if err != nil {
				x.Stop()
			}
		}()

		if opts == nil {
			err = errors.New("Invalid options")
			return
		}

		if err = opts.parse(); err != nil {
			return
		}

		x.opts = opts

		if x.isListening() {
			err = errors.New("Server already running.")
			return
		}

		err = x.execBoot()

		if err != nil {
			return
		}

		x.opts.freeupAddress()

		x.listener, err = net.Listen(opts.network, opts.address)

		if err != nil {
			return
		}
	})

	return
}

func (x *Server) Run() (err error) {
	graceful.Run(func(grace graceful.Grace) {
		if x.listener == nil {
			err = errors.New("Server already running.")
			return
		}

		defer x.Stop()

		if !x.isListening() {
			err = errors.New("Server not running.")
			return
		}

		for {
			select {
			case <-grace.Done():
				return

			case resp := <-x.listener.AcceptChan():
				if resp.Error != nil {
					return
				}

				inbound := &Inbound{
					conn:     resp.Data,
					updQueue: make(map[string]chan string),
					server:   x,
				}

				go inbound.Run()
			}
		}
	})

	return
}

func (x *Server) execBoot() (err error) {
	for _, callback := range x.bootCallbacks {
		err = callback()

		if err != nil {
			break
		}
	}

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

func (x *Server) forEach(callback InboundPredicate) {
	x.mutex.Lock()

	defer x.mutex.Unlock()

	for _, inbound := range x.mapper {
		err := callback(inbound)

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
