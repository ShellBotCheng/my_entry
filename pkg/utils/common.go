package utils

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
)

const RandomLen int =  32

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GenerateRandomString 获取随机字符串
func GenerateRandomString(len int) (string, error) {
	b, err := GenerateRandomBytes(len)
	return base64.URLEncoding.EncodeToString(b), err
}

func MD5(s string) string {
	h := md5.New()
	//uid不会发生变化
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}


