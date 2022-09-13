package mapping

import (
	"ahutoj/web/dao"
	"ahutoj/web/io/request"
)

func ProblemReqToDao(req request.Problem) dao.Problem {
	problem := dao.Problem{Pid: req.Pid}

	// 	LimitMemory:   req.LimitMemory,
	if req.Title != nil {
		problem.Title = *req.Title
	}
	if req.Description != nil {
		problem.Description = *req.Description
	}
	if req.Input != nil {
		problem.Input = *req.Input
	}
	if req.Output != nil {
		problem.Output = *req.Output
	}
	if req.Sample_input != nil {
		problem.Sample_input = *req.Sample_input
	}
	if req.Sample_output != nil {
		problem.Sample_output = *req.Sample_output
	}
	if req.Hit != nil {
		problem.Hit = *req.Hit
	}
	if req.Label != nil {
		problem.Label = *req.Label
	}
	if req.LimitTime != nil {
		problem.LimitTime = *req.LimitTime
	}
	if req.LimitMemory != nil {
		problem.LimitMemory = *req.LimitMemory
	}
	return problem
}
