package crypto

import "github.com/Cyberpull/gokit/errors"

type xKey struct {
	value string
}

func (x *xKey) Get(key ...string) (value string, err error) {
	if len(key) > 0 {
		value = key[0]
		return
	}

	if x.value != "" {
		value = x.value
		return
	}

	err = errors.New("Cipher Key is required")

	return
}

// ========================

var Key xKey

func SetDefaultKey(key string) {
	Key.value = key
}
