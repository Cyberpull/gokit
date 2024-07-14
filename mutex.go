package gokit

import "sync"

type Mutex[T comparable] struct {
	mutex  sync.Mutex
	mapper map[T]*xMutexEntry[T]
}

func (x *Mutex[T]) Lock(key T) {
	x.mutex.Lock()

	defer x.mutex.Unlock()

	x.of(key).Lock()
}

func (x *Mutex[T]) Unlock(key T) {
	x.mutex.Lock()

	defer x.mutex.Unlock()

	x.of(key).Unlock()
}

func (x *Mutex[T]) of(key T) *xMutexEntry[T] {
	entry, ok := d(x).mapper[key]

	if !ok {
		entry = &xMutexEntry[T]{parent: x, key: key}
		x.mapper[key] = entry
	}

	return entry
}

func (x *Mutex[T]) remove(key T) {
	delete(x.mapper, key)
}

func (x *Mutex[T]) initialize() {
	if x.mapper == nil {
		x.mapper = make(map[T]*xMutexEntry[T])
	}
}
