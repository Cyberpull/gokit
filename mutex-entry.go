package gokit

import "sync"

type xMutexEntry[T comparable] struct {
	parent *Mutex[T]
	mutex  sync.Mutex
	count  int
	key    T
}

func (x *xMutexEntry[T]) Lock() {
	x.count++
	x.mutex.Lock()
}

func (x *xMutexEntry[T]) Unlock() {
	x.mutex.Unlock()
	x.count--

	if x.count <= 0 {
		x.parent.remove(x.key)
	}
}
