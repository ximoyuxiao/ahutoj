package request

type List struct {
	LID       int64  `json:"LID"`
	UID       string `json:"UID"`
	Title     string `json:"Title"`
	StartTime int64  `json:"StartTime"`
}

type ListProblem struct {
	LID int64 `json:"LID"`
	PID int   `json:"PID"`
}

type ListAll struct {
	LID       int64  `json:"LID"`
	UID       string `json:"UID"`
	PID       int    `json:"PID"`
	Title     string `json:"Title"`
	StartTime int64  `json:"StartTime"`
}
type TrainingListReq GetListReq
