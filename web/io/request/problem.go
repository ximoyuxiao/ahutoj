package request

type Problem struct {
	PID           int64   `json:"PID"`
	Title         *string `json:"Title"`
	Description   *string `json:"Description"`
	Input         *string `json:"Input"`
	Output        *string `json:"Output"`
	Sample_input  *string `json:"SampleInput"`
	Sample_output *string `json:"SampleOutput"`
	LimitTime     *int64  `json:"LimitTime"`
	LimitMemory   *int64  `json:"LimitMemory"`
	Hit           *string `json:"Hit"`
	Label         *string `json:"Label"`
	Origin        *int64  `json:"Origin"`
	OriginPID     *string `json:"OriginPID"`
	ContentType   int64   `json:"ContentType"`
	Visible       int     `json:"Visible"`
}
type EditProblemReq Problem
type AddProblemReq Problem
type DeleteProblemReq struct {
	PIDs []int64 `json:"PIDs"`
}

type ProblemListReq GetListReq
