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
	Default  ModuleCode = 100
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
	VerifyError               OperationCode = 37 //用户权限不足
	DataEmpty                 OperationCode = 40 //数据为空
	DataNotExist              OperationCode = 41 //数据不存在
	DataResolutionError       OperationCode = 42 //数据解析失败
	UIDExist                  OperationCode = 50 //UID已存在
	UIDNotExist               OperationCode = 51 //UID不存在
	UIDEmpty                  OperationCode = 52 //UID为空
	PasswordError             OperationCode = 53 //密码错误
	PasswordEmpty             OperationCode = 54 //密码为空
	PIDNotExist               OperationCode = 55 //PID不存在
	CIDNotExist               OperationCode = 56 //CID不存在
	CIDPassWordError          OperationCode = 57 //CID密码错误
	ContestNotBegin           OperationCode = 60 //竞赛未开始
	FileUnsupport             OperationCode = 71 //文件不支持
	Duplicate                 OperationCode = 72 //副本
	Notimplemented            OperationCode = 99 //接口未实现
)

func GetResCode(mod ModuleCode, loc LocationCode, op OperationCode) ResCode {
	res := int32(int(mod)*1000 + int(loc)*100 + int(op))
	return ResCode(res)
}

/*状态码  做一个规范*/
//老状态码  未来要删除
const (
	ParametersInvlidCode ResCode = 100132
	ServerBusyCode       ResCode = 100131
)

var codeMsgMap = map[ResCode]string{
	SuccessCode:          "success",
	ParametersInvlidCode: "参数错误",
	ServerBusyCode:       "服务器错误",
}
var HttpCodeMap = map[ResCode]int{
	SuccessCode:          http.StatusOK,
	ParametersInvlidCode: http.StatusBadRequest,
	ServerBusyCode:       http.StatusInternalServerError,
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
