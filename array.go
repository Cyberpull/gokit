package gokit

import "sync"

type ArrayCallback[T comparable] func(entry T, i int)
type ArrayPredicate[T comparable] func(entry T, i int) bool

type Array[T comparable] struct {
	mutex   sync.Mutex
	entries []T
}

func (x *Array[T]) Append(entries ...T) {
	x.mutex.Lock()

	defer x.mutex.Unlock()

	d(x).entries = append(d(x).entries, entries...)
}

func (x *Array[T]) Prepend(entries ...T) {
	x.mutex.Lock()

	defer x.mutex.Unlock()

	d(x).entries = append(entries, d(x).entries...)
}

func (x *Array[T]) Filter(callback ArrayPredicate[T]) (value *Array[T]) {
	x.mutex.Lock()

	defer x.mutex.Unlock()

	value = &Array[T]{entries: make([]T, 0)}

	for i, entry := range d(x).entries {
		valid := callback(entry, i)

		if valid {
			value.entries = append(value.entries, entry)
		}
	}

	return
}

func (x *Array[T]) ForEach(callback ArrayCallback[T]) {
	x.mutex.Lock()

	defer x.mutex.Unlock()

	for i, entry := range d(x).entries {
		callback(entry, i)
	}
}

func (x *Array[T]) Contains(v T) bool {
	x.mutex.Lock()

	defer x.mutex.Unlock()

	for _, entry := range x.entries {
		if entry == v {
			return true
		}
	}

	return false
}

func (x *Array[T]) IndexOf(v T) int {
	x.mutex.Lock()

	defer x.mutex.Unlock()

	for i, entry := range x.entries {
		if entry == v {
			return i
		}
	}

	return -1
}

func (x *Array[T]) LastIndexOf(v T) int {
	x.mutex.Lock()

	defer x.mutex.Unlock()

	lastIndex := len(d(x).entries) - 1

	for i := lastIndex; i >= 0; i-- {
		entry := x.entries[i]

		if entry == v {
			return i
		}
	}

	return -1
}

func (x *Array[T]) TakeAt(i int) T {
	if i >= x.Len() {
		panic("Index out of range")
	}

	x.mutex.Lock()

	defer x.mutex.Unlock()

	entry := d(x).entries[i]

	x.entries = append(x.entries[:i], x.entries[i+1:]...)

	return entry
}

func (x *Array[T]) TakeFirst() T {
	return x.TakeAt(0)
}

func (x *Array[T]) TakeLast() T {
	return x.TakeAt(x.Len() - 1)
}

func (x *Array[T]) Len() int {
	x.mutex.Lock()

	defer x.mutex.Unlock()

	return len(x.entries)
}

func (x *Array[T]) Slice() []T {
	x.mutex.Lock()

	defer x.mutex.Unlock()

	return d(x).entries
}

func (x *Array[T]) Clear() {
	x.mutex.Lock()

	defer x.mutex.Unlock()

	x.entries = make([]T, 0)
}

func (x *Array[T]) initialize() {
	if x.entries == nil {
		x.entries = make([]T, 0)
	}
}
