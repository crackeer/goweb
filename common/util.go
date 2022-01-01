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

func RandString(len int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}
