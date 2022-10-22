package constanct

import (
	"net/http"
)

// ResCode int32
type ResCode int32

//新状态码 定义
//模块码
type ModuleCode int32

//位置码
type LocationCode int32

//操作码
type OperationCode int32

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

//新状态码
const SuccessCode ResCode = 0
const (
	Auth     ModuleCode = 101
	User     ModuleCode = 102
	Admin    ModuleCode = 103
	Problem  ModuleCode = 104
	Training ModuleCode = 105
	Contest  ModuleCode = 106
	Submit   ModuleCode = 106
	File     ModuleCode = 107
)
const (
	Service LocationCode = 1
	Logic   LocationCode = 2
	Models  LocationCode = 3
)
const (
	MysqlAdd                  OperationCode = 11
	MysqlDelete               OperationCode = 12
	MysqlUpdate               OperationCode = 13
	MysqlQuery                OperationCode = 14
	RedisAdd                  OperationCode = 15
	RedisDelete               OperationCode = 16
	RedisUpdate               OperationCode = 17
	RedisQuery                OperationCode = 18
	ServerBusy                OperationCode = 30 //服务器繁忙
	ServerError               OperationCode = 31 //服务器错误
	Parsesparameters          OperationCode = 32 //解析参数
	ParametersTypeError       OperationCode = 33 //参数类型错误
	ParametersFormatError     OperationCode = 34 //参数格式错误
	ParametersConversionError OperationCode = 35 //参数转换失败
	TokenBuildError           OperationCode = 36 //Token创建错误
	DataEmpty                 OperationCode = 40 //数据为空
	DataNotExist              OperationCode = 41 //数据不存在
	DataResolutionError       OperationCode = 42 //数据解析失败
	UIDExist                  OperationCode = 50 //UID已存在
	UIDNotExist               OperationCode = 51 //UID不存在
	UIDEmpty                  OperationCode = 52 //UID为空
	PasswordError             OperationCode = 53 //密码错误
	PasswordEmpty             OperationCode = 54 //密码为空
	Notimplemented            OperationCode = 99 //接口未实现
)

func GetResCode(mod ModuleCode, loc LocationCode, op OperationCode) ResCode {
	res := int32(int(mod)*1000 + int(loc)*100 + int(op))
	return ResCode(res)
}

/*状态码  做一个规范*/
//老状态码  未来要删除
const (
	PassEmpty            ResCode = 102
	PageNotFound         ResCode = 404
	VerifyErrorCode      ResCode = 501
	InvalidParamCode     ResCode = 1001
	UIDNotExistCode      ResCode = 1002
	NotLoginCode         ResCode = 1003
	PassWordErrorCode    ResCode = 1004
	TokenBuildErrorCode  ResCode = 1005
	TokenInvalidCode     ResCode = 1006
	UIDExistCOde         ResCode = 1007
	PIDExistCode         ResCode = 1008
	PIDNotExistCode      ResCode = 1009
	CIDNotExistCode      ResCode = 1010
	CIDPassWordErrorCode ResCode = 2000
	CONTESTNOTEBEGIN     ResCode = 2005
	MySQLErrorCode       ResCode = 2001
	RedisErrorCode       ResCode = 2002
	DUPLICATECODE        ResCode = 2003
	ServerBusyCode       ResCode = 5001
	FILEUNSUPPORT        ResCode = 6001
)

var codeMsgMap = map[ResCode]string{
	SuccessCode:          "success",
	PageNotFound:         "页面未找到",
	PassEmpty:            "密码为空",
	InvalidParamCode:     "请求参数错误",
	UIDNotExistCode:      "账号不存在",
	NotLoginCode:         "账号未登录",
	TokenBuildErrorCode:  "Token创建失败",
	TokenInvalidCode:     "无效的Token",
	PassWordErrorCode:    "密码错误",
	MySQLErrorCode:       "数据库错误",
	RedisErrorCode:       "缓存数据库错误",
	ServerBusyCode:       "服务器繁忙",
	UIDExistCOde:         "该用户已存在",
	PIDExistCode:         "该题目已存在",
	CIDNotExistCode:      "竞赛不存在",
	PIDNotExistCode:      "题目不存在",
	VerifyErrorCode:      "用户权限不足",
	CIDPassWordErrorCode: "竞赛密码错误",
	FILEUNSUPPORT:        "不支持的文件类型",
	CONTESTNOTEBEGIN:     "竞赛未开始",
}
var HttpCodeMap = map[ResCode]int{
	SuccessCode: http.StatusOK,
	// UIDEmpty:            http.StatusOK,
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
	CIDNotExistCode:     http.StatusOK,
	PIDNotExistCode:     http.StatusOK,
	VerifyErrorCode:     http.StatusForbidden,
	FILEUNSUPPORT:       http.StatusOK,
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
