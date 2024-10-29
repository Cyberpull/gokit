package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"

	"github.com/Cyberpull/gokit/errors"
)

type xDecrypt struct{}

func (x *xDecrypt) AES(encrypted string, cipherKey ...string) (value string, err error) {
	key, err := Key.Get(cipherKey...)
	if err != nil {
		return
	}

	cipherText, err := hex.DecodeString(encrypted)

	if err != nil {
		return
	}

	block, err := aes.NewCipher([]byte(key))

	if err != nil {
		return
	}

	if len(cipherText) < aes.BlockSize {
		err = errors.New("cipherText block size is too short")
		return
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	value = string(cipherText)

	return
}

// ======================

var Decrypt xDecrypt
