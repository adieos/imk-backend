package utils

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

func Sha256Encrypt(s string) string {
	hash := sha256.New()
	hash.Write([]byte(s))
	return hex.EncodeToString(hash.Sum(nil))
}

func Sha512Encrypt(s string) string {
	hash := sha512.New()
	hash.Write([]byte(s))
	return hex.EncodeToString(hash.Sum(nil))
}
