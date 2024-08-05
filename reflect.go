package gokit

import (
	"reflect"

	"github.com/Cyberpull/gokit/errors"
)

func Write(input any, output any) (err error) {
	switch data := input.(type) {
	case []byte:
		return write(data, output)

	case string:
		return write([]byte(data), output)

	default:
		return errors.New("Invalid input")
	}
}

func New[T any]() (value T) {
	vType := reflect.TypeOf(value)

	switch vType.Kind() {
	case reflect.Pointer:
		value = reflect.New(vType.Elem()).Interface().(T)

	default:
		value = reflect.Zero(vType).Interface().(T)
	}

	return
}

func Zero[T any]() (value T) {
	vType := reflect.TypeOf(value)
	value = reflect.Zero(vType).Interface().(T)
	return
}
