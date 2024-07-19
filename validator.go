package gokit

import (
	"reflect"

	vengine "github.com/go-playground/validator/v10"
)

type Validator interface {
	Validate(data any, rules ...any) (err error)
}

// ==========================

type validator struct {
	core *vengine.Validate
}

func (v *validator) Validate(data any, rules ...any) (err error) {
	rValue := reflect.ValueOf(data)

	rType := rValue.Type()
	rKind := rType.Kind()

	if rKind == reflect.Pointer {
		rValue = rValue.Elem()
	}

	value := rValue.Interface()

	switch rKind {
	case reflect.Struct:
		return v.core.Struct(value)
	default:
		rule := oneOfAny[string]("", rules)
		return v.core.Var(value, rule)
	}
}

// ==========================

const defaultTagName string = "validate"

func NewValidator(tagName ...string) Validator {
	tag := one(defaultTagName, tagName)

	core := vengine.New()
	core.SetTagName(tag)

	return &validator{
		core: core,
	}
}
