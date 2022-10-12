package response

import "ahutoj/web/io/constanct"

type GetSubmitResp struct {
	Response
	SID        int                `json:"SID"`
	UID        string             `json:"UID"`
	PID        int                `json:"PID"`
	Source     string             `json:"Source"`
	Lang       constanct.LANG     `json:"Lang"`
	Result     constanct.OJResult `json:"Result"`
	UseTime    int64              `json:"UseTime"`
	UseMemory  int64              `json:"UseMemory"`
	SubmitTime int64              `json:"SubmitTime"`
	CeInfo     *string            `json:"CeInfo"`
}
type SubmitLIstItem struct {
	SID        int                `json:"SID"`
	UID        string             `json:"UID"`
	PID        int                `json:"PID"`
	Lang       constanct.LANG     `json:"Lang"`
	Result     constanct.OJResult `json:"Result"`
	UseTime    int64              `json:"UseTime"`
	UseMemory  int64              `json:"UseMemory"`
	SubmitTime int64              `json:"SubmitTime"`
}
type SubmitListResp struct {
	Response
	Count    int64            `json:"Count"`
	LastTime int64            `json:"LastTime"`
	Data     []SubmitLIstItem `json:"Data"`
}
