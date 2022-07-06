package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"

	"github.com/sanity-io/litter"
)

func MD5EnCode(userID, password string) (string, error) {
	Md5 := md5.New()
	var t int32 = 0
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
	return hex.EncodeToString(Md5.Sum(nil)), nil
}

func Sdump(value ...interface{}) string {
	return litter.Sdump(value...)
}
