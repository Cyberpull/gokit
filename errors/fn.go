package errors

import (
	"fmt"
	"runtime/debug"
)

func errorCode(code ...int) int {
	if len(code) == 0 {
		code = append(code, 500)
	}

	if code[0] == 0 || code[0] == 200 {
		code[0] = 500
	}

	return code[0]
}

func New(message string, code ...int) (err *Error) {
	return &Error{
		code:    errorCode(code...),
		message: message,
	}
}

func WithStack(message string, code ...int) (err *Error) {
	err = New(message, code...)

	defer recover()

	if stack := debug.Stack(); len(stack) > 0 {
		err.stack = string(stack)
	}

	return err
}

func Newf(message string, v ...any) *Error {
	code := make([]int, 0)

	if len(v) > 0 {
		vcode, ok := v[0].(int)

		if ok {
			code = append(code, vcode)
		}

		v = v[1:]
	}

	message = fmt.Sprintf(message, v...)

	return New(message, code...)
}

func From(v any, code ...int) *Error {
	var value *Error

	switch x := v.(type) {
	case *Error:
		value = x

		if len(code) > 0 {
			value.code = errorCode(code...)
		}
	case string:
		value = New(x, code...)
	case error:
		value = New(x.Error(), code...)
	default:
		value = New("An unknown error occurred", 500)
	}

	return value
}
