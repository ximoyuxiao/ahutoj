package response

import "ahutoj/web/dao"

type ProblemItemResp struct {
	PID   int64  `json:"PID"`
	Title string `json:"Title"`
	Label string `json:"Label"`
}

type ProblemListResp struct {
	Response
	Count int64             `json:"Count"`
	Data  []ProblemItemResp `json:"Data"`
}

type ProblemResp dao.Problem

type ProblemInfoResp struct {
	Response
	ProblemResp
}
