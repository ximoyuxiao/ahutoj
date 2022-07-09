package request

type Problem struct {
	Pid           int    `form:"pid"`
	Title         string `form:"title"`
	Description   string `form:"description"`
	Input         string `form:"input"`
	Output        string `form:"output"`
	Sample_input  string `form:"sample_input"`
	Sample_output string `form:"sample_output"`
	Hit           string `form:"hit"`
	LimitTime     int    `form:"limitTime"`
	LimitMemory   int    `form:"limitMemory"`
}
