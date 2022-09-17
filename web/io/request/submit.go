package request

import "ahutoj/web/io/constanct"

type AddSubmitReq struct {
	PID        int            `json:"PID"`
	UID        string         `json:"UID"`
	CID        int            `json:"CID"`
	Source     string         `json:"Source"`
	Lang       constanct.LANG `json:"Lang"`
	SubmitTime int64          `json:"SubmitTime"`
}

type RejudgeSubmitReq struct {
	SID *int    `json:"SID"`
	PID *int    `json:"PID"`
	UID *string `json:"UID"`
	CID *int    `json:"CID"`
}

type GetSubmitReq struct {
	SID int64 `param:"SID"`
}

type SubmitListReq struct {
	PID *int    `query:"PID"`
	UID *string `query:"UID"`
	CID *int    `query:"CID"`
	GetListReq
}

type GetCodeReq struct {
	SID int `json:"SID"`
}
