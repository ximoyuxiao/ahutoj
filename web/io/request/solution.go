package request

// Request
type SolutionReq struct {
	ActionType int64  `json:"ActionType"` // 1创建，2删除
	Pid        string `json:"PID"`        // 题目ID
	Sid        string `json:"SID"`
	Text       string `json:"Text,omitempty"`  // 文本内容
	Title      string `json:"Title,omitempty"` // 标题
	Uid        string `json:"UID"`             // 用户ID
}
type SolutionListReq struct {
	PID string `json:"PID"` // 题目ID
}
