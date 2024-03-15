package response

import "ahutoj/web/dao"

type ProblemItemResp struct {
	PID   string `json:"PID"`
	Title string `json:"Title"`
	Label string `json:"Label"`
	Accepted int64 `json:"Accepted"`
	Submited int64  `json:"Submit"`
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
	SolutionNumber int64 `json:"SolutionNumber"`
}
type AddProblemResp struct {
	Response
	PID string `json:"PID"`
}
