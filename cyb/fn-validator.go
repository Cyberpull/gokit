package cyb

import "github.com/Cyberpull/gokit"

type xValidator struct {
	engine gokit.Validator
}

func (x xValidator) Validate(data any, rules ...any) (err error) {
	return x.engine.Validate(data, rules...)
}

// ===============================

var validator xValidator

func init() {
	validator.engine = gokit.NewValidator("binding")
}

func validate(data any, rules ...any) (err error) {
	return validator.Validate(data, rules...)
}
