package request

type CreateNoticeReq struct {
	Content string `json:"Content"`
	Title   string `json:"Title"`
}

type UpdateNoticeReq struct {
	ID      int    `param:"ID"`
	Content string `json:"Content"`
	Title   string `json:"Title"`
}

type GetNoticeReq struct {
	ID int `param:"ID"`
}

type DeleteNoticeReq GetNoticeReq
