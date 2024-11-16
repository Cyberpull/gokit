package scopes

func e[T any](v T) T {
	switch e := any(v).(type) {
	case Column:
		if e.Table == "" {
			e.Table = CurrentTable
		}

	case Table:
		if e.Name == "" {
			e.Name = CurrentTable
		}
	}

	return v
}

// Expressions
func ex[T Expression](e []T) (v []Expression) {
	v = make([]Expression, 0)

	for _, x := range e {
		v = append(v, x)
	}

	return
}
