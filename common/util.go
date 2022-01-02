package common

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"time"
)

func GetNowTimeString() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func GetNowDate() string {
	return time.Now().Format("2006-01-02")
}

func Md5Crypt(data []byte) string {
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has) //将[]byte转成16进制
}

var defaultLetters = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

// RandomString returns a random string with a fixed length
func RandomString(n int) string {
	letters := defaultLetters

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
