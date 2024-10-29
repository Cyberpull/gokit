package gokit

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/Cyberpull/gokit/errors"

	"golang.org/x/exp/constraints"
)

type JoinEntryFunc func(v string) string

type SetConstraint interface {
	constraints.Integer | constraints.Float | ~string
}

func In[T comparable](value T, entries ...T) bool {
	for _, entry := range entries {
		if value == entry {
			return true
		}
	}

	return false
}

func JoinFunc[T any](delim string, entries []T, callbacks ...JoinEntryFunc) string {
	var buff bytes.Buffer

	for _, entry := range entries {
		data := fmt.Sprint(entry)
		data = strings.TrimSpace(data)

		for _, callback := range callbacks {
			if callback == nil {
				continue
			}

			data = callback(data)
		}

		if data == "" {
			continue
		}

		if buff.Len() > 0 {
			buff.WriteString(delim)
		}

		buff.WriteString(data)
	}

	return buff.String()
}

func Join[T any](delim string, entries ...T) string {
	return JoinFunc(delim, entries)
}

func Split[T SetConstraint](data string, delim string) (value []T, err error) {
	var t T

	value = make([]T, 0)

	for _, entry := range strings.Split(data, delim) {
		var newValue any

		switch any(t).(type) {
		case string:
			newValue = entry

		case int, int8, int16, int32, int64:
			newValue, err = strconv.ParseInt(entry, 0, 64)

		case uint, uint8, uint16, uint32, uint64:
			newValue, err = strconv.ParseUint(entry, 0, 64)

		case float32, float64:
			newValue, err = strconv.ParseFloat(entry, 64)

		default:
			err = errors.Newf("Type '%T' not allowed", t)
		}

		if err != nil {
			return
		}

		value = append(value, any(newValue).(T))
	}

	return
}
