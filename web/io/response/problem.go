package response

import "ahutoj/web/dao"

type ProblemItemResp struct {
	Pid   int    `json:"pid"`
	Title string `json:"title"`
}

type ProblemListResp struct {
	Response
	Count int               `json:"count"`
	Data  []ProblemItemResp `json:"data"`
}

type ProblemResp dao.Problem

type ProblemInfoResp struct {
	Response
	ProblemResp
}
