package utils

import (
	"ahutoj/web/io/constanct"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"math/rand"
	"os/exec"
	"time"

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

func MD5EnCodeStr(str string) (string, error) {
	Md5 := md5.New()
	io.WriteString(Md5, str)
	return hex.EncodeToString(Md5.Sum(nil)), nil
}

func Sdump(value ...interface{}) string {
	return litter.Sdump(value...)
}

func GetPageInfo(page, limit int) (int, int) {
	retSize := constanct.GetDefaultLimit()
	if limit != 0 {
		retSize = limit
	}
	offset := constanct.GetDefaultOffset()
	if page > offset {
		offset = retSize * page
	}
	return offset, retSize
}

func MinInt64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func MaxInt64(a, b int64) int64 {
	if a < b {
		return b
	}
	return a
}

func CheckNum(ch int8) bool {
	return ch <= '9' && ch >= '0'
}

func ExecuteCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	err := cmd.Start()
	return err
}
func GetNow() int64 {
	return time.Now().UnixMilli()
}

func GenVeriey(size int) string {
	data := "1234567890qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"
	res := ""
	for i := 0; i < size; i++ {
		number := rand.Int() % len(data)
		res += string(data[number])
	}
	return res
}
