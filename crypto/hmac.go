package crypto

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
)

type xHashAlgo func() hash.Hash

type xHmac struct{}

func (x *xHmac) Hash(algo xHashAlgo, key []byte, data []byte) (value string, err error) {
	hash := hmac.New(algo, key)

	if _, err = hash.Write(data); err != nil {
		return
	}

	value = hex.EncodeToString(hash.Sum(nil))

	return
}

func (x *xHmac) MD5(key []byte, data []byte) (value string, err error) {
	return x.Hash(md5.New, key, data)
}

func (x *xHmac) Sha1(key []byte, data []byte) (value string, err error) {
	return x.Hash(sha1.New, key, data)
}

func (x *xHmac) Sha256(key []byte, data []byte) (value string, err error) {
	return x.Hash(sha256.New, key, data)
}

func (x *xHmac) Sha512(key []byte, data []byte) (value string, err error) {
	return x.Hash(sha512.New, key, data)
}

// ==================

var HMAC xHmac
