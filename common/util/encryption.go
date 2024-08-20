package util

import (
	"crypto/md5"
	"encoding/hex"
)

func EncryptonMd5(s string) string {
	hash := md5.New()
	hash.Write([]byte(s))
	return hex.EncodeToString(hash.Sum(nil))
}
