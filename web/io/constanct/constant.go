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
	SuccessCode ResCode = 0
)

/*公共错误 10*/
const (
	ServerErrorCode    ResCode = 100001
	InvalidParamCode   ResCode = 100002
	PageNotFoundCode   ResCode = 100003
	NotimplementedCode ResCode = 100004
)

/*auth 11*/
const (
	AUTH_Token_EmptyCode     ResCode = 110001
	AUTH_Token_InvalidCode   ResCode = 110002
	AUTH_Token_URLVerifyCode ResCode = 110003 // 用户无访问接口权限

	AUTH_LOGIN_FAILED          ResCode = 110101
	AUTH_LOGIN_UIDEmptyCode    ResCode = 110102
	AUTH_LOGIN_PassEmptyCode   ResCode = 110103
	AUTH_LOGIN_PASSERRORCode   ResCode = 110104
	AUTH_LOGIN_UIDNotExistCode ResCode = 110105
	AUTH_LOGIN_TokenBuildCode  ResCode = 110106

	AUTH_REGISTER_FAILED         ResCode = 110201
	AUTH_REGISTER_UIDExistCode   ResCode = 110202
	AUTH_REGISTER_TokenBuildCode ResCode = 110203
)

/*user 12*/
const (
	USER_INFO_FAILED          ResCode = 120101
	USER_INFO_UIDNotExistCode ResCode = 120102

	USER_EDITINFO_FAILED          ResCode = 120201
	USER_EDITINFO_UIDNotExistCode ResCode = 120202

	USER_EDITPASS_FAILED            ResCode = 120301
	USER_EDITPASS_PasswordEmptyCode ResCode = 120302
	USER_EDITPASS_PasswordCode      ResCode = 120303

	USER_STATUS_FAILED ResCode = 120401

	USER_CFBIND_FAILED        ResCode = 120501
	USER_CFBIND_UserEmptyCode ResCode = 120502
	USER_CFBIND_PassEmptyCode ResCode = 120503
	USER_CFBIND_PassErrorCode ResCode = 120504
)

/*admin 13*/
const (
	ADMIN_ADD_FAILED   ResCode = 130101
	ADMIN_ADD_UIDEmpty ResCode = 130102

	ADMIN_EDIT_FAILED ResCode = 130201
	ADMIN_EDIT_ADMIN  ResCode = 130202

	ADMIN_DELETE_FAILED ResCode = 130301

	ADMIN_LIST_FAILED ResCode = 130401

	ADMIN_GET_FAILED ResCode = 130501
)

/*problem 14*/
const (
	PROBLEM_ADD_FAILED ResCode = 140101

	PROBLEM_EDIT_FAILED           ResCode = 140201
	PROBLEM_EDIT_PIDNoteExistCode ResCode = 140202

	PROBLEM_DELETE_FAILED ResCode = 140301

	PROBLEM_LIST_FAILED ResCode = 140401

	PROBLEM_GET_FAILED          ResCode = 140501
	PROBLEM_GET_PIDNotExistCode ResCode = 140502
)

/*traning 15*/
const (
	TRAIN_ADD_FAILED ResCode = 150101

	TRAIN_EDIT_FAILED ResCode = 150201

	TRAIN_DELETE_FAILED ResCode = 150301

	TRAIN_LIST_FAILED ResCode = 150401

	TRAIN_GET_FAILED ResCode = 150501
)

/*contest 16*/
const (
	CONTEST_ADD_FAILED ResCode = 160101

	CONTEST_EDIT_FAILED ResCode = 160201

	CONTEST_DELETE_FAILED ResCode = 160301

	CONTEST_LIST_FAILED ResCode = 160401

	CONTEST_GET_FAILED               ResCode = 160501
	CONTEST_GET_CIDNotExistCode      ResCode = 160502
	CONTEST_GET_NotBegin             ResCode = 160503
	CONTEST_GET_CIDPassWordErrorCode ResCode = 160504

	CONTEST_RANK_FAILED ResCode = 160601
)

/*submit 17*/
const (
	SUBMIT_ADD_FAILEDCode    ResCode = 170101
	SUBMIT_ADD_DUPLICATECODE ResCode = 170102

	SUBMIT_REJUDG_FAILEDCode ResCode = 170201

	SUBMIT_LIST_FAILEDCode ResCode = 170401

	SUBMIT_GET_FAILEDCode ResCode = 170501
)

/*file 18*/
const (
	FILE_UP_FAILEDCode    ResCode = 180101
	FILE_UP_UNSUPPORTCode ResCode = 180102

	FILE_REMOVE_FAILEDCode ResCode = 180201

	FILE_UNZIP_FAILEDCode    ResCode = 180301
	FILE_UNZIP_UNSUPPORTCode ResCode = 180302
	FILE_UNZIP_NotExistCode  ResCode = 180303

	FILE_LIST_FAILEDCode ResCode = 180401
)

var codeMsgMap = map[ResCode]string{
	SuccessCode:                      "",
	InvalidParamCode:                 "请求参数错误",
	ServerErrorCode:                  "服务器错误",
	PageNotFoundCode:                 "页面不存在",
	NotimplementedCode:               "接口未实现",
	AUTH_Token_EmptyCode:             "用户未登录",
	AUTH_Token_InvalidCode:           "用户登录信息过期",
	AUTH_Token_URLVerifyCode:         "用户没有权限访问",
	AUTH_LOGIN_FAILED:                "登陆失败",
	AUTH_LOGIN_UIDEmptyCode:          "用户ID不能为空",
	AUTH_LOGIN_PassEmptyCode:         "用户密码不能为空",
	AUTH_LOGIN_UIDNotExistCode:       "用户不存在",
	AUTH_LOGIN_PASSERRORCode:         "密码错误",
	AUTH_LOGIN_TokenBuildCode:        "登陆失败",
	AUTH_REGISTER_FAILED:             "注册失败",
	AUTH_REGISTER_UIDExistCode:       "用户ID已存在",
	AUTH_REGISTER_TokenBuildCode:     "注册失败",
	USER_INFO_FAILED:                 "用户信息获取失败",
	USER_INFO_UIDNotExistCode:        "用户ID不存在",
	USER_EDITINFO_FAILED:             "用户信息保存失败",
	USER_EDITINFO_UIDNotExistCode:    "用户ID不存在",
	USER_EDITPASS_FAILED:             "修改密码失败",
	USER_EDITPASS_PasswordEmptyCode:  "修改密码不能为空",
	USER_EDITPASS_PasswordCode:       "请输入正确的原密码",
	USER_STATUS_FAILED:               "用户信息获取失败",
	USER_CFBIND_FAILED:               "绑定codeforce账号失败",
	USER_CFBIND_UserEmptyCode:        "codeforce账号不能为空",
	USER_CFBIND_PassEmptyCode:        "codeforce密码不能为空",
	USER_CFBIND_PassErrorCode:        "codeforce密码错误",
	ADMIN_ADD_FAILED:                 "添加权限失败",
	ADMIN_ADD_UIDEmpty:               "待添加用户ID不能为空",
	ADMIN_EDIT_FAILED:                "修改用户权限失败",
	ADMIN_EDIT_ADMIN:                 "不能修改admin的权限信息",
	ADMIN_DELETE_FAILED:              "删除用户权限失败",
	ADMIN_LIST_FAILED:                "获取用户权限列表失败",
	ADMIN_GET_FAILED:                 "获取用户权限信息失败",
	PROBLEM_ADD_FAILED:               "添加题目失败",
	PROBLEM_EDIT_FAILED:              "编辑题目失败",
	PROBLEM_EDIT_PIDNoteExistCode:    "题目不存在",
	PROBLEM_DELETE_FAILED:            "删除题目失败",
	PROBLEM_LIST_FAILED:              "获取题目列表失败",
	PROBLEM_GET_FAILED:               "获取题目信息失败",
	PROBLEM_GET_PIDNotExistCode:      "题目不存在",
	TRAIN_ADD_FAILED:                 "添加题单失败",
	TRAIN_EDIT_FAILED:                "编辑题单失败",
	TRAIN_DELETE_FAILED:              "删除题单失败",
	TRAIN_LIST_FAILED:                "获取提单列表失败",
	TRAIN_GET_FAILED:                 "获取题单信息失败",
	CONTEST_ADD_FAILED:               "添加竞赛失败",
	CONTEST_EDIT_FAILED:              "编辑竞赛失败",
	CONTEST_DELETE_FAILED:            "删除竞赛失败",
	CONTEST_LIST_FAILED:              "获取竞赛列表失败",
	CONTEST_GET_FAILED:               "获取竞赛信息失败",
	CONTEST_GET_CIDNotExistCode:      "该竞赛不存在",
	CONTEST_GET_NotBegin:             "该竞赛未开始",
	CONTEST_GET_CIDPassWordErrorCode: "竞赛密码错误",
	CONTEST_RANK_FAILED:              "获取竞赛排名失败",
	SUBMIT_ADD_FAILEDCode:            "提交代码失败",
	SUBMIT_ADD_DUPLICATECODE:         "不能重复提交代码",
	SUBMIT_REJUDG_FAILEDCode:         "重判题目失败",
	SUBMIT_LIST_FAILEDCode:           "获取提交状态信息失败",
	SUBMIT_GET_FAILEDCode:            "获取提交信息失败",
	FILE_UP_FAILEDCode:               "上传文件失败",
	FILE_UP_UNSUPPORTCode:            "文件类型不支持",
	FILE_REMOVE_FAILEDCode:           "删除文件失败",
	FILE_UNZIP_FAILEDCode:            "解压文件失败",
	FILE_UNZIP_UNSUPPORTCode:         "文件类型不支持",
	FILE_UNZIP_NotExistCode:          "不存在的文件",
	FILE_LIST_FAILEDCode:             "获取判题文件失败",
}
var HttpCodeMap = map[ResCode]int{
	SuccessCode:        http.StatusOK,
	InvalidParamCode:   http.StatusBadRequest,
	ServerErrorCode:    http.StatusServiceUnavailable,
	PageNotFoundCode:   http.StatusNotFound,
	NotimplementedCode: http.StatusForbidden,
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[ServerErrorCode]
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
