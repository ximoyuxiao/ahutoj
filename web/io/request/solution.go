package request

// Request
type SolutionReq struct {
	ActionType int64  `json:"ActionType"`      // 1创建，2编辑,3删除
	Pid        string `json:"PID"`             // 题目ID
	Sid        int64  `json:"SID"`             // 题解
	Text       string `json:"Text,omitempty"`  // 文本内容
	Title      string `json:"Title,omitempty"` // 标题
	Uid        string `json:"UID,omitempty"`   // 用户ID
}
type GetSolutionReq struct {
	SID int64 `json:"SID"`
}

type GetSolutionListReq struct {
	PID string `query:"PID"`
	UID string `query:"UID"`
	GetListReq
}

type EditSolutionReq SolutionReq
type CreateSolutionReq SolutionReq
