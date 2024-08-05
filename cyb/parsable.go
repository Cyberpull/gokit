package cyb

import (
	"bytes"
	"fmt"

	"github.com/Cyberpull/gokit/errors"
)

type parsable interface {
	name() string
	prefix() string
}

func hasPrefix[T parsable](x T, data []byte) bool {
	return bytes.HasPrefix(data, []byte(x.prefix()))
}

func parse[T parsable](x T, data []byte) (err error) {
	prefix := []byte(x.prefix())

	if !bytes.HasPrefix(data, prefix) {
		err = errors.New(fmt.Sprintf("Unable to parse %v.", x.name()))
		return
	}

	data = bytes.TrimPrefix(data, prefix)

	err = parseJson(data, x)

	return
}

func toBytes[T parsable](x T) (value []byte, err error) {
	value, err = toJson(x)

	if err != nil {
		return
	}

	value = append([]byte(x.prefix()), value...)

	return
}

// func toString[T parsable](x T) (value string, err error) {
// 	data, err := toBytes(x)

// 	if err != nil {
// 		return
// 	}

// 	value = string(data)

// 	return
// }
