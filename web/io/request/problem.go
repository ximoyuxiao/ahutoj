package request

type Problem struct {
	Pid           int     `json:"pid"`
	Title         *string `json:"title"`
	Description   *string `json:"description"`
	Input         *string `json:"input"`
	Output        *string `json:"output"`
	Sample_input  *string `json:"sample_input"`
	Sample_output *string `json:"sample_output"`
	LimitTime     *int64  `json:"limitTime"`
	LimitMemory   *int64  `json:"limitMemory"`
	Hit           *string `json:"hit"`
	Label         *string `json:"label"`
}
type EditProblemReq Problem
type AddProblemReq Problem
type DeleteProblemReq struct {
	Pids []int64 `json:"pids"`
}

type ProblemListReq GetListReq
