package request

type List struct {
	Lid   int64  `json:"lid"`
	Uid   string `json:"uid"`
	Title string `json:"title"`
	Stime int64  `json:"stime"`
}

type ListProblem struct {
	Lid int64 `json:"lid"`
	Pid int   `json:"pid"`
}

type ListAll struct {
	Lid   int64  `json:"lid"`
	Uid   string `json:"uid"`
	Pid   int    `json:"pid"`
	Title string `json:"title"`
	Stime int64  `json:"stime"`
}
