package response

type GetNOticeItem struct {
	ID          int    `json:"ID"`
	UID         string `json:"UID"`
	Title       string `json:"Title"`
	Content     string `json:"Content"`
	CreatedTime int64  `json:"CreatedTime"`
	UpdatedTime int64  `json:"UpdatedTime"`
}
type GetNoticeResp struct {
	Response
	GetNOticeItem
}

type GetListNoticeResp struct {
	Response
	Count int             `json:"Count"`
	Data  []GetNOticeItem `json:"data"`
}
