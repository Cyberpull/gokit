package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
)

type xEncrypt struct{}

func (x *xEncrypt) AES(text string, cipherKey ...string) (value string, err error) {
	key, err := Key.Get(cipherKey...)

	if err != nil {
		return
	}

	block, err := aes.NewCipher([]byte(key))

	if err != nil {
		return
	}

	cipherText := make([]byte, aes.BlockSize+len(text))
	iv := cipherText[:aes.BlockSize]

	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], []byte(text))

	value = hex.EncodeToString(cipherText)

	return
}

// ======================

var Encrypt xEncrypt
