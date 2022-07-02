package utils

import (
	"crypto/md5"
	"fmt"
	"io"
)

func MD5EnCode(userID, password string) (string, error) {
	Md5 := md5.New()
	t := 0
	len := len(password) % 8
	if len == 0 {
		len = 8
	}
	k := 32 - len
	for k > 0 {
		t = t | (1 << (k))
		k -= len

	}
	code := fmt.Sprintf("%s&%d&%s", userID, t, password)
	io.WriteString(Md5, code)
	return string(Md5.Sum(nil)), nil
}
