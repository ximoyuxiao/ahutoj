package request

import "ahutoj/web/io/constanct"

type AddSubmitReq struct {
	Pid        int            `json:"pid"`
	Uid        string         `json:"uid"`
	Cid        int            `json:"cid"`
	Source     string         `json:"source"`
	Lang       constanct.LANG `json:"lang"`
	SubmitTime int64          `json:"submitTime"`
}

type RejudgeSubmitReq struct {
	Sid *int    `json:"sid"`
	Pid *int    `json:"Pid"`
	Uid *string `json:"Uid"`
	Cid *int    `json:"Cid"`
}

type GetSubmitReq struct {
	Sid int64 `param:"sid"`
}

type SubmitListReq struct {
	Pid *int    `query:"Pid"`
	Uid *string `query:"Uid"`
	Cid *int    `query:"Cid"`
	GetListReq
}

type GetCodeReq struct {
	Sid int `json:"sid"`
}
