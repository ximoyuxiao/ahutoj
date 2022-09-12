package constanct

import "net/http"

// ResCode int32
type ResCode int32

const (
	DefaultLimit  int = 20
	DefaultOffset int = 0
)

func GetDefaultLimit() int {
	return DefaultLimit
}

func GetDefaultOffset() int {
	return DefaultOffset
}

const (
	SuccessCode         ResCode = 0
	UIDEmpty            ResCode = 101
	PassEmpty           ResCode = 102
	PageNotFound        ResCode = 404
	VerifyErrorCode     ResCode = 501
	InvalidParamCode    ResCode = 1001
	UIDNotExistCode     ResCode = 1002
	NotLoginCode        ResCode = 1003
	PassWordErrorCode   ResCode = 1004
	TokenBuildErrorCode ResCode = 1005
	TokenInvalidCode    ResCode = 1006
	UIDExistCOde        ResCode = 1007
	PIDExistCode        ResCode = 1008
	PIDNotExistCode     ResCode = 1009
	MySQLErrorCode      ResCode = 2001
	RedisErrorCode      ResCode = 2002
	ServerBusyCode      ResCode = 5001
	Notimplemented      ResCode = 9999
)

var codeMsgMap = map[ResCode]string{
	SuccessCode:         "success",
	PageNotFound:        "页面未找到",
	UIDEmpty:            "账号为空",
	PassEmpty:           "密码为空",
	InvalidParamCode:    "请求参数错误",
	UIDNotExistCode:     "账号不存在",
	NotLoginCode:        "账号未登录",
	TokenBuildErrorCode: "Token创建失败",
	TokenInvalidCode:    "无效的Token",
	PassWordErrorCode:   "密码错误",
	MySQLErrorCode:      "数据库错误",
	RedisErrorCode:      "缓存数据库错误",
	ServerBusyCode:      "服务器繁忙",
	UIDExistCOde:        "该用户已存在",
	PIDExistCode:        "该题目已存在",
	PIDNotExistCode:     "题目不存在",
	VerifyErrorCode:     "用户权限不足",
	Notimplemented:      "接口未实现",
}
var HttpCodeMap = map[ResCode]int{
	SuccessCode:         http.StatusOK,
	UIDEmpty:            http.StatusOK,
	PassEmpty:           http.StatusOK,
	PageNotFound:        http.StatusNotFound,
	InvalidParamCode:    http.StatusBadRequest,
	UIDNotExistCode:     http.StatusOK,
	NotLoginCode:        http.StatusOK,
	TokenBuildErrorCode: http.StatusUnauthorized,
	TokenInvalidCode:    http.StatusUnauthorized,
	PassWordErrorCode:   http.StatusOK,
	MySQLErrorCode:      http.StatusInternalServerError,
	RedisErrorCode:      http.StatusInternalServerError,
	ServerBusyCode:      http.StatusInternalServerError,
	UIDExistCOde:        http.StatusOK,
	PIDExistCode:        http.StatusOK,
	PIDNotExistCode:     http.StatusOK,
	VerifyErrorCode:     http.StatusForbidden,
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[ServerBusyCode]
	}
	return msg
}

func (c ResCode) HttpCode() int {
	HttpCode, ok := HttpCodeMap[c]
	if !ok {
		HttpCode = http.StatusOK
	}
	return HttpCode
}
