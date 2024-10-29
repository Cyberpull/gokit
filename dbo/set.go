package dbo

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/Cyberpull/gokit"
)

type Set[T gokit.SetConstraint] struct {
	Data []T
}

func (x *Set[T]) Scan(value any) (err error) {
	switch data := value.(type) {
	case []T:
		x.Data = data

	default:
		x.Data, err = gokit.Split[T](fmt.Sprint(value), ",")
	}

	return
}

func (x Set[T]) Value() (value driver.Value, err error) {
	value = gokit.Join(",", x.Data...)
	return
}

func (x *Set[T]) UnmarshalJSON(b []byte) error {
	var value T

	if err := json.Unmarshal(b, &value); err != nil {
		return err
	}

	return x.Scan(value)
}

func (x Set[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.Data)
}
