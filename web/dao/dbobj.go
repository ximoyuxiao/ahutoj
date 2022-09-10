package dao

type User struct {
	Uid     string `db:"uid"`
	Uname   string `db:"uname"`
	Pass    string `db:"pass"`
	School  string `db:"school"`
	Classes string `db:"classes"`
	Major   string `db:"major"`
	Adept   string `gorm:"column:adept"`
	Vjid    string `db:"vjid"`
	Vjpwd   string `db:"vjpwd"`
	Email   string `db:"email"`
}

func (u User) TableName() string {
	return "User"
}

type Permission struct {
	Uid             string `db:"uid"`
	Administrator   string `db:"administrator"`
	Problem_edit    string `db:"problem_edit"`
	Source_browser  string `db:"source_browser"`
	Contest_creator string `db:"contest_creator"`
}

func (p Permission) TableName() string {
	return "Permission"
}

type Problem struct {
	Pid           int    `db:"pid" json:"Pid"`
	Title         string `db:"title" json:"title"`
	Description   string `db:"description" json:"description"`
	Input         string `db:"input" json:"input"`
	Output        string `db:"output" json:"output"`
	Sample_input  string `db:"sample_input" json:"sample_input"`
	Sample_output string `db:"sample_output" json:"sample_output"`
	LimitTime     int64  `gorm:"column:limitTime" json:"limitTime"`
	LimitMemory   int64  `gorm:"column:limitMemory" json:"limitMemory"`
	Hit           string `db:"hit" json:"hit"`
	Label         string `db:"label" json:"label"`
}

type List struct {
	Lid   int64  `db:"lid"`
	Uid   string `db:"uid"`
	Title string `db:"title"`
	Stime int64  `db:"stime"`
}

type ListProblem struct {
	Lid int64 `db:"lid"`
	Pid int   `db:"pid"`
}

type ListUser struct {
	Lid        int64 `db:"lid"`
	Uid        int   `db:"uid"`
	Submit_num int   `db:"submit_num"`
	Ac_num     int   `db:"ac_num"`
}

type Contest struct {
	Cid         int64  `db:"cid"`
	Uid         string `db:"uid"`
	Title       string `db:"title"`
	Description string `db:"description"`
	Begin_time  int64  `db:"begin_time"`
	End_time    int64  `db:"end_time"`
	Ctype       int    `db:"ctype"`
	Ispublic    string `db:"ispublic"`
	Pass        string `db:"pass"`
}

func (p Contest) TableName() string {
	return "Contest"
}

type ConPro struct {
	Cid        int64  `db:"cid"`
	Pid        int    `db:"pid"`
	Ptitle     string `db:"title"`
	Submit_num int    `db:"submit_num"`
	Ac_num     int    `db:"ac_num"`
}

func (p ConPro) TableName() string {
	return "Conpro"
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

func (p Submit) TableName() string {
	return "Submit"
}
