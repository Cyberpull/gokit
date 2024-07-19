package gokit

import (
	"encoding/json"
	"reflect"

	"cyberpull.com/gokit/errors"
)

func one[T any](def T, attr []T) T {
	if len(attr) > 0 {
		return attr[0]
	}

	return def
}

func oneOfAny[T any](def T, attr []any) T {
	if len(attr) > 0 {
		return attr[0].(T)
	}

	return def
}

func write(data []byte, output any) (err error) {
	defer func() {
		r := recover()

		if r != nil {
			err = errors.From(r)
		}
	}()

	if output == nil {
		err = errors.New("Null output is not allowed", 500)
		return
	}

	vtype := reflect.TypeOf(output)

	if vtype.Kind() != reflect.Pointer {
		err = errors.New("Only pointer output is allowed", 500)
		return
	}

	if err = json.Unmarshal(data, output); err != nil {
		err = nil

		value := reflect.ValueOf(output).Elem()

		switch value.Kind() {
		case reflect.String:
			value.Set(reflect.ValueOf(string(data)))
			return

		default:
			switch value.Interface().(type) {
			case []byte:
				value.Set(reflect.ValueOf(data))
				return
			}
		}

		err = errors.New("Unable to parse data", 500)
	}

	return
}
