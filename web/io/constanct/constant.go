package constanct

// ResCode int32
type ResCode int32

const (
	SuccessCode         ResCode = 0
	InvalidParamCode    ResCode = 1001
	UIDNotExistCode     ResCode = 1002
	NotLoginCode        ResCode = 1003
	PassWordErrorCode   ResCode = 1004
	TokenBuildErrorCode ResCode = 1005
	TokenInvaildCode    ResCode = 1006
	UIDExistCOde        ResCode = 1007
	PIDExistCode        ResCode = 1008
	MySQLErrorCode      ResCode = 2001
	RedisErrorCode      ResCode = 2002
	ServerBusyCode      ResCode = 5001
)

var codeMsgMap = map[ResCode]string{
	SuccessCode:         "success",
	InvalidParamCode:    "请求参数错误",
	UIDNotExistCode:     "账号不存在",
	NotLoginCode:        "账号未登录",
	TokenBuildErrorCode: "Token创建失败",
	TokenInvaildCode:    "无效的Token",
	PassWordErrorCode:   "密码错误",
	MySQLErrorCode:      "数据库错误",
	RedisErrorCode:      "缓存数据库错误",
	ServerBusyCode:      "服务器繁忙",
	UIDExistCOde:        "该用户已存在",
	PIDExistCode:        "该题目已存在",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[ServerBusyCode]
	}
	return msg
}
