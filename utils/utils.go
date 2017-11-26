package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(src string) string {
	signByte := []byte(src)
	hash := md5.New()
	hash.Write(signByte)
	return hex.EncodeToString(hash.Sum(nil))
}
