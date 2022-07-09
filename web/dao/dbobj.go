package dao

type User struct {
	Uid     string `db:"uid"`
	Uname   string `db:"uname"`
	Pass    string `db:"pass"`
	School  string `db:"school"`
	Classes string `db:"classes"`
	Major   string `db:"major"`
	Vjid    string `db:"vjid"`
	Vjpwd   string `db:"vjpwd"`
	Email   string `db:"email"`
}

type Permission struct {
	Uid             string `db:"uid"`
	Administrator   string `db:"administrator"`
	Problem_edit    string `db:"problem_edit"`
	Source_browser  string `db:"source_browser"`
	Contest_creator string `db:"contest_creator"`
}

type Problem struct {
	Pid           *int   `db:"pid"`
	Title         string `db:"title"`
	Description   string `db:"description"`
	Input         string `db:"input"`
	Output        string `db:"output"`
	Sample_input  string `db:"sample_input"`
	Sample_output string `db:"sample_output"`
	Hit           string `db:"hit"`
	LimitTime     int    `db:"limitTime"`
	LimitMemory   int    `db:"limitMemory"`
}

type List struct {
	Lid   int    `db:"lid"`
	Uid   string `db:"uid"`
	Title string `db:"title"`
	Stime int64  `db:"stime"`
}

type ListProblem struct {
	Lid int `db:"lid"`
	Pid int `db:"pid"`
}

type ListUser struct {
	Lid        int `db:"lid"`
	Uid        int `db:"uid"`
	Submit_num int `db:"submit_num"`
	Ac_num     int `db:"ac_num"`
}

type Contest struct {
	Cid         int    `db:"cid"`
	Uid         string `db:"uid"`
	Title       string `db:"title"`
	Description string `db:"description"`
	Begin_time  int64  `db:"begin_time"`
	End_time    int64  `db:"end_time"`
	Ctype       string `db:"ctype"`
	Ispublic    string `db:"ispublic"`
	Pass        string `db:"pass"`
}

type ConPro struct {
	Cid        int `db:"cid"`
	Pid        int `db:"pid"`
	Submit_num int `db:"submit_num"`
	Ac_num     int `db:"ac_num"`
}

type Submit struct {
	Sid        int    `db:"sid"`
	Pid        int    `db:"pid"`
	Uid        string `db:"uid"`
	Cid        int    `db:"cid"`
	Judgeid    int    `db:"judgeid"`
	Source     string `db:"source"`
	Lang       string `db:"lang"`
	Result     string `db:"result"`
	Usetime    int    `db:"usetime"`
	Memory     int    `db:"memory"`
	SubmitTime int64  `db:"submittime"`
}
