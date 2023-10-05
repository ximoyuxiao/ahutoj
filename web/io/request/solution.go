package request

// Request
type SolutionReq struct {
	ActionType int64  `json:"ActionType"`      // 1创建，2编辑,3删除
	Pid        string `json:"PID"`             // 题目ID
	Sid        string `json:"SID"`             // 题解
	Text       string `json:"Text,omitempty"`  // 文本内容
	Title      string `json:"Title,omitempty"` // 标题
	Uid        string `json:"UID,omitempty"`   // 用户ID
}
type SolutionListReq struct {
	PID string `json:"PID"` // 题目ID
}

type GetSolutionReq struct {
	SID int `json:"SID"`
}

type GetSolutionListReq struct {
	PID string `json:"PID"`
	UID string `json:"UID"`
}

type EditSolutionReq SolutionReq
type CreateSolutionReq SolutionReq
