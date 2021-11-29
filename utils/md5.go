package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMD5StringOfBytes(bytes []byte) string {
	hash := md5.Sum(bytes)
	return hex.EncodeToString(hash[:])
}

func GetMD5StringOfString(text string) string {
	return GetMD5StringOfBytes([]byte(text))
}
