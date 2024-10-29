package crypto

import (
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"math"
)

type xRand struct{}

func (x *xRand) Bytes(max int) (value []byte, err error) {
	value = make([]byte, max)

	if _, err = rand.Read(value); err != nil {
		return
	}

	return
}

func (x *xRand) Uint64(max int) (value uint64, err error) {
	b, err := x.Bytes(max)

	if err != nil {
		return
	}

	value = binary.BigEndian.Uint64(b)

	return
}

func (x *xRand) Int(max int) (value int, err error) {
	result, err := x.Uint64(max)

	if err != nil {
		return
	}

	value = int(result)

	return
}

func (x *xRand) Float64(max int) (value float64, err error) {
	result, err := x.Uint64(max)

	if err != nil {
		return
	}

	value = float64(result)

	return
}

func (x *xRand) String(max int) (value string, err error) {
	max = int(math.Ceil(float64(max) / 2))

	b, err := x.Bytes(max)

	if err != nil {
		return
	}

	value = hex.EncodeToString(b)

	return
}

// ==================

var Rand xRand

// Deprecated: Use `Rand.Bytes()` instead
func RandomBytes(max int) (value []byte, err error) {
	value = make([]byte, max)

	if _, err = rand.Read(value); err != nil {
		return
	}

	return
}

// Deprecated: Use `Rand.Int()` instead
func RandomInt(max int) (value int, err error) {
	b, err := RandomBytes(max)

	if err != nil {
		return
	}

	value = int(binary.BigEndian.Uint64(b))

	return
}

// Deprecated: Use `Rand.String()` instead
func RandomString(max int) (value string, err error) {
	max = int(math.Ceil(float64(max) / 2))
	b, err := RandomBytes(max)

	if err != nil {
		return
	}

	value = hex.EncodeToString(b)

	return
}
