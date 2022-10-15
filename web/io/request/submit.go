package request

import "ahutoj/web/io/constanct"

type AddSubmitReq struct {
	PID        int64          `json:"PID"`
	UID        string         `json:"UID"`
	CID        int64          `json:"CID"`
	Source     string         `json:"Source"`
	Lang       constanct.LANG `json:"Lang"`
	SubmitTime int64          `json:"SubmitTime"`
}

type RejudgeSubmitReq struct {
	SID *int64  `json:"SID"`
	PID *int64  `json:"PID"`
	UID *string `json:"UID"`
	CID *int64  `json:"CID"`
}

type GetSubmitReq struct {
	SID int64 `param:"SID"`
}

type SubmitListReq struct {
	PID *int64  `query:"PID"`
	UID *string `query:"UID"`
	CID *int64  `query:"CID"`
	GetListReq
}

type GetCodeReq struct {
	SID int64 `json:"SID"`
}
