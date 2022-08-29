package request

type List struct {
	Lid   int64  `json:"lid"`
	Uid   string `json:"uid"`
	Title string `json:"title"`
	Stime int64  `json:"stime"`
}

type ListProblem struct {
	Lid int `json:"lid"`
	Pid int `json:"pid"`
}
