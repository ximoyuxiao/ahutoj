package response

import "ahutoj/web/io/constanct"

type GetSubmitResp struct {
	Response
	Sid        int                `json:"sid"`
	Pid        int                `json:"pid"`
	Source     string             `json:"Source"`
	Lang       constanct.LANG     `json:"lang"`
	Result     constanct.OJResult `json:"result"`
	UseTime    int                `json:"useTime"`
	UseMemory  int                `json:"useMemory"`
	SubmitTime int64              `json:"submitTime"`
}
type SubmitLIstItem struct {
	Sid        int                `json:"sid"`
	Pid        int                `json:"pid"`
	Lang       constanct.LANG     `json:"lang"`
	Result     constanct.OJResult `json:"result"`
	UseTime    int                `json:"useTime"`
	UseMemory  int                `json:"useMemory"`
	SubmitTime int64              `json:"submitTime"`
}
type SubmitListResp struct {
	Response
	Count    int64            `json:"count"`
	LastTime int64            `json:"lastTime"`
	Data     []SubmitLIstItem `json:"Data"`
}
