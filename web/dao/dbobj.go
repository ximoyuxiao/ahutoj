package dao

import "ahutoj/web/io/constanct"

type User struct {
	Uid     string `gorm:"column:uid"`
	Uname   string `gorm:"column:uname"`
	Pass    string `gorm:"column:pass"`
	School  string `gorm:"column:school"`
	Classes string `gorm:"column:classes"`
	Major   string `gorm:"column:major"`
	Adept   string `gorm:"column:adept"`
	Vjid    string `gorm:"column:vjid"`
	Vjpwd   string `gorm:"column:vjpwd"`
	Email   string `gorm:"column:email"`
}

func (u User) TableName() string {
	return "User"
}

type Permission struct {
	Uid             string `gorm:"column:uid"`
	Administrator   string `gorm:"column:administrator"`
	Problem_edit    string `gorm:"column:problem_edit"`
	Source_browser  string `gorm:"column:source_browser"`
	Contest_creator string `gorm:"column:contest_creator"`
}

func (p Permission) TableName() string {
	return "Permission"
}

type Problem struct {
	Pid           int    `gorm:"column:pid" json:"pid"`
	Title         string `gorm:"column:title" json:"title"`
	Description   string `gorm:"column:description" json:"description"`
	Input         string `gorm:"column:input" json:"input"`
	Output        string `gorm:"column:output" json:"output"`
	Sample_input  string `gorm:"column:sample_input" json:"sample_input"`
	Sample_output string `gorm:"column:sample_output" json:"sample_output"`
	LimitTime     int64  `gorm:"column:limit_time" json:"limitTime"`
	LimitMemory   int64  `gorm:"column:limit_memory" json:"limitMemory"`
	Hit           string `gorm:"column:hit" json:"hit"`
	Label         string `gorm:"column:label" json:"label"`
}

type List struct {
	Lid   int64  `gorm:"column:lid"`
	Uid   string `gorm:"column:uid"`
	Title string `gorm:"column:title"`
	Stime int64  `gorm:"column:stime"`
}

type ListProblem struct {
	Lid int64 `gorm:"column:lid"`
	Pid int   `gorm:"column:pid"`
}

type ListUser struct {
	Lid        int64 `gorm:"column:lid"`
	Uid        int   `gorm:"column:uid"`
	Submit_num int   `gorm:"column:submit_num"`
	Ac_num     int   `gorm:"column:ac_num"`
}

type Contest struct {
	Cid         int64  `gorm:"column:cid"`
	Uid         string `gorm:"column:uid"`
	Title       string `gorm:"column:title"`
	Description string `gorm:"column:description"`
	Begin_time  int64  `gorm:"column:begin_time"`
	End_time    int64  `gorm:"column:end_time"`
	Ctype       int    `gorm:"column:ctype"`
	Ispublic    int    `gorm:"column:ispublic"`
	Pass        string `gorm:"column:pass"`
}

func (p Contest) TableName() string {
	return "Contest"
}

type ConPro struct {
	Cid        int64  `gorm:"column:cid"`
	Pid        int    `gorm:"column:pid"`
	Ptitle     string `gorm:"column:ptitle"`
	Submit_num int    `gorm:"column:submit_num"`
	Ac_num     int    `gorm:"column:ac_num"`
}

func (p ConPro) TableName() string {
	return "ConPro"
}

type Submit struct {
	Sid        int                `gorm:"column:sid"`
	Pid        int                `gorm:"column:pid"`
	Uid        string             `gorm:"column:uid"`
	Cid        int                `gorm:"column:cid"`
	Judgeid    int                `gorm:"column:judgeid"`
	Source     string             `gorm:"column:source"`
	Lang       constanct.LANG     `gorm:"column:lang"`
	Result     constanct.OJResult `gorm:"column:result"`
	Usetime    int                `gorm:"column:usetime"`
	Memory     int                `gorm:"column:memory"`
	SubmitTime int64              `gorm:"column:submittime"`
}

func (p Submit) TableName() string {
	return "Submit"
}
