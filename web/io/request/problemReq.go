package request

type Problem struct {
	Pid           int    `json:"pid"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	Input         string `json:"input"`
	Output        string `json:"output"`
	Sample_input  string `json:"sample_input"`
	Sample_output string `json:"sample_output"`
	Hit           string `json:"hit"`
	LimitTime     int    `json:"limitTime"`
	LimitMemory   int    `json:"limitMemory"`
}
