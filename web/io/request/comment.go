package request

type CommentReq struct {
	ActionType int64  `json:"ActionType"`      // 1创建,2删除
	Sid        int64  `json:"SID"`             // 题解
	Text       string `json:"Text,omitempty"`  // 文本内容
	Title      string `json:"Title,omitempty"` // 标题
	Uid        string `json:"UID,omitempty"`   // 用户ID
	CID        int64  `json:"CID"`
	FCID       int64  `json:"FCID"`
}

type CommentListReq struct {
	SID int64 `query:"SID"`
	GetListReq
}
