package request

// Request
type FavoriteReq struct {
	ActionType int64  `json:"ActionType"`
	SID        int64  `json:"SID"` // 点赞题解
	UID        string `json:"UID"`
}
