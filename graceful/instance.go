package graceful

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type xGraceful struct {
	mutex   sync.Mutex
	counter int64
	ctx     context.Context
	cancel  context.CancelFunc
	signal  chan os.Signal
	exiting bool
}

func (x *xGraceful) Run(callback Runner) {
	if x.isExiting() {
		return
	}

	d(x).add()

	defer d(x).remove()

	callback(x.ctx)
}

func (x *xGraceful) Start() {
	x.mutex.Lock()

	defer x.mutex.Unlock()

	if x.signal != nil {
		return
	}

	x.signal = make(chan os.Signal, 1)

	signal.Notify(x.signal, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-x.signal

		fmt.Println("\nShutting down...")

		d(x).exit()

		for {
			if x.count() == 0 {
				fmt.Println("Goodbye!")
				os.Exit(0)
			}
		}
	}()
}

// ========================

func (x *xGraceful) add() {
	x.mutex.Lock()

	defer x.mutex.Unlock()

	d(x).counter++
}

func (x *xGraceful) remove() {
	x.mutex.Lock()

	defer x.mutex.Unlock()

	x.counter--

	if x.counter < 0 {
		x.counter = 0
	}
}

func (x *xGraceful) count() int64 {
	x.mutex.Lock()

	defer x.mutex.Unlock()

	return x.counter
}

func (x *xGraceful) exit() {
	x.mutex.Lock()

	defer x.mutex.Unlock()

	x.exiting = true

	d(x).cancel()
}

func (x *xGraceful) isExiting() bool {
	x.mutex.Lock()

	defer x.mutex.Unlock()

	return x.exiting
}

func (x *xGraceful) initialize() {
	if x.ctx == nil {
		x.ctx = context.TODO()
		x.ctx, x.cancel = context.WithCancel(x.ctx)
	}
}

// ========================

var instance xGraceful

func init() {
	instance.Start()
}
