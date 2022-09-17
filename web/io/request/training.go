package request

type List struct {
	LID   int64  `json:"LID"`
	UID   string `json:"UID"`
	Title string `json:"Title"`
	Stime int64  `json:"Stime"`
}

type ListProblem struct {
	LID int64 `json:"LID"`
	PID int   `json:"PID"`
}

type ListAll struct {
	LID   int64  `json:"LID"`
	UID   string `json:"UID"`
	PID   int    `json:"PID"`
	Title string `json:"Title"`
	Stime int64  `json:"StartTime"`
}
