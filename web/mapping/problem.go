package mapping

import (
	"ahutoj/web/dao"
	"ahutoj/web/io/request"
)

func ProblemReqToDao(req request.Problem) dao.Problem {
	problem := dao.Problem{}
	if req.PID != nil {
		problem.PID = *req.PID
	}
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
		problem.SampleInput = *req.Sample_input
	}
	if req.Sample_output != nil {
		problem.SampleOutput = *req.Sample_output
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
	if req.Origin != nil {
		problem.Origin = *req.Origin
	}
	if req.OriginPID != nil {
		problem.OriginPID = *req.OriginPID
	}
	problem.PType = req.PType
	problem.ContentType = req.ContentType
	problem.Visible = req.Visible
	return problem
}
