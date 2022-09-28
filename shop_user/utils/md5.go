package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func PWMd5(password string) string {
	h := md5.New()
	md5code := "42gh25jytfa"
	h.Write([]byte(password + md5code))
	return hex.EncodeToString(h.Sum(nil))
}
