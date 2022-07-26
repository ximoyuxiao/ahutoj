package response

import "ahutoj/web/dao"

type ProblemItemResp struct {
	Pid   int    `json:"pid"`
	Title string `json:"title"`
}

type ProblemListResp struct {
	Response
	Data []ProblemItemResp `json:"Data"`
}

type ProblemResp dao.Problem

type ProblemInfoResp struct {
	Response
	ProblemResp
}
