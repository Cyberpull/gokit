package crypto

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

type xHash struct{}

func (x *xHash) Hash(algo xHashAlgo, data []byte) (value string, err error) {
	hash := algo()

	if _, err = hash.Write(data); err != nil {
		return
	}

	value = hex.EncodeToString(hash.Sum(nil))

	return
}

func (x *xHash) MD5(data []byte) (value string, err error) {
	return x.Hash(md5.New, data)
}

func (x *xHash) Sha1(data []byte) (value string, err error) {
	return x.Hash(sha1.New, data)
}

func (x *xHash) Sha256(data []byte) (value string, err error) {
	return x.Hash(sha256.New, data)
}

func (x *xHash) Sha512(data []byte) (value string, err error) {
	return x.Hash(sha512.New, data)
}

// ==================

var Hash xHash
