package utils

import (
	"compress/gzip"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type HttpMethodType string
type UserStatus int64

const (
	POST   HttpMethodType = "POST"
	GET    HttpMethodType = "GET"
	DELETE HttpMethodType = "DELETE"
	PUT    HttpMethodType = "PUT"
)

/*useRedirect 是否使用重定向，不是的话跳过重定向，是就不管*/
func DoRequest(method HttpMethodType, url string, headers map[string]string, cookies map[string]string, body *string, useRedirect bool) (*http.Response, error) {
	logger := GetLogInstance()
	client := &http.Client{}
	if !useRedirect {
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
	}
	var data io.Reader = nil
	contntLength := 0
	if body != nil {
		data = strings.NewReader(*body)
		contntLength = len(*body)
	}
	req, err := http.NewRequest(string(method), url, data)
	if err != nil {
		logger.Errorf("call NewRequest failed,method=%v, url=%v, data=%v, err=%v", method, url, body, err.Error())
		return nil, err
	}
	req.ContentLength = int64(contntLength)
	if cookies != nil {
		req.Header.Set("Cookie", MapToStrings(cookies, ";"))
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	resp, err := client.Do(req)
	if err != nil {
		logger.Errorf("call Do failed,req=%+v ,err=%s", Sdump(req), err.Error())
		return nil, err
	}
	return resp, nil
}

func MapToStrings(data map[string]string, sep string) string {
	if data == nil {
		return ""
	}
	strs := make([]string, 0)
	for key, value := range data {
		strs = append(strs, key+"="+value)
	}
	return strings.Join(strs, sep)
}

func MapToFormStrings(data map[string]string, sep string) string {
	if data == nil {
		return ""
	}
	strs := make([]string, 0)
	for key, value := range data {
		key = url.QueryEscape(key)
		value = url.QueryEscape(value)
		strs = append(strs, key+"="+value)
	}
	return strings.Join(strs, sep)
}

func ParseRespToReader(resp *http.Response) io.ReadCloser {
	var reader io.ReadCloser
	if resp.Header.Get("Content-Encoding") == "gzip" {
		reader, _ = gzip.NewReader(resp.Body)
	} else {
		reader = resp.Body
	}
	return reader
}

func ParseRespToByte(resp *http.Response) ([]byte, error) {
	var reader io.ReadCloser
	if resp.Header.Get("Content-Encoding") == "gzip" {
		reader, _ = gzip.NewReader(resp.Body)
	} else {
		reader = resp.Body
	}
	return io.ReadAll(reader)
}

func DealStrings(str string) string {
	ret := ""
	for _, ch := range str {
		if ch == '\n' || ch == '\t' {
			continue
		}
		ret += string(ch)
	}
	return strings.Trim(ret, " ")
}
