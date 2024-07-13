package yaml

import (
	"reflect"

	"gopkg.in/yaml.v3"
)

func Encode(data any) (value []byte, err error) {
	return yaml.Marshal(data)
}

func Decode[T any](b []byte) (value T, err error) {
	vType := reflect.TypeOf(value)

	switch vType.Kind() {
	case reflect.Pointer:
		value = reflect.New(vType.Elem()).Interface().(T)
		err = yaml.Unmarshal(b, value)
	default:
		err = yaml.Unmarshal(b, &value)
	}

	return
}
