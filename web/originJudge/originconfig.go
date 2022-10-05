package originjudge

import "ahutoj/web/dao"

type OJPlatform int64
type OriginFunc interface {
	Judge()
	login()
	submit()
	getResult()
}
type OriginJudge struct {
	PID    string     // 平台的题目ID
	Submit dao.Submit // 用户提交代码
}
