package utils

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"
	"time"
)

func CheckPathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func GetFileSuffix(filename string) string {
	strs := strings.Split(filename, ".")
	return strs[len(strs)-1]
}

func ChekfileHashSuffix(filename string, args ...string) bool {
	suffix := GetFileSuffix(filename)
	for _, arg := range args {
		if arg == suffix {
			return true
		}
	}
	return false
}

func SaveFile(savePath string, filename string, data []byte) error {
	ok, err := CheckPathExists(savePath)
	if err != nil {
		return err
	}
	if !ok {
		err = os.Mkdir(savePath, 0777)
		if err != nil {
			return err
		}
	}
	targetFile := savePath + "/" + filename
	os.WriteFile(targetFile, data, 0777)
	return nil
}

func EncodeBase64FromByte(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func EncodeBase64FromString(data string) string {
	return EncodeBase64FromByte([]byte(data))
}

func DeCodeBase64ToByte(data string) ([]byte, error) {
	res, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func DeCodeBase64ToString(data string) (string, error) {
	res, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

func GetFileName(filename string, bytes []byte) string {
	now := time.Now().UnixNano()
	strs := strings.Split(filename, ".")
	md5str, _ := MD5EnCodeStr(string(bytes))
	suffix := strs[len(strs)-1]
	imageName := fmt.Sprintf("%v%v.%v", md5str, now, suffix)
	return imageName
}
