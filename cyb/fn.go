package cyb

import (
	"reflect"

	"github.com/Cyberpull/gokit/errors"
)

func MakeRequest[T any](client *Client, method, channel string, data any) (value T, err error) {
	if client == nil {
		err = errors.New("Null client supplied.")
		return
	}

	resp, err := client.Request(method, channel, data)

	if err != nil {
		return
	}

	vType := reflect.TypeOf(value)

	switch vType.Kind() {
	case reflect.Pointer:
		value = reflect.New(vType.Elem()).Interface().(T)
		err = resp.Bind(value)

	default:
		err = resp.Bind(&value)
	}

	return
}
