package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(content string) string {
	h := md5.New()
	h.Write([]byte(content))
	return hex.EncodeToString(h.Sum(nil))
}
