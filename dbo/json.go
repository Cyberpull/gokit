package dbo

import (
	"database/sql/driver"
	"encoding/json"

	"cyberpull.com/gotk/v2/errors"
)

type Json[T any] struct {
	Data T
}

func (n *Json[T]) Scan(value any) (err error) {
	switch data := value.(type) {
	case T:
		n.Data = data

	case string:
		err = n.UnmarshalJSON([]byte(data))

	case []byte:
		err = n.UnmarshalJSON(data)

	default:
		err = errors.Newf("Data type '%T' not supported", value)
	}

	return
}

func (n Json[T]) Value() (value driver.Value, err error) {
	return n.MarshalJSON()
}

func (n *Json[T]) UnmarshalJSON(b []byte) error {
	var value T

	if err := json.Unmarshal(b, &value); err != nil {
		return err
	}

	return n.Scan(value)
}

func (n Json[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.Data)
}
