package dao

import "ahutoj/web/io/constanct"

type User struct {
	UID           string `gorm:"column:UID"`
	Uname         string `gorm:"column:UserName"`
	Pass          string `gorm:"column:Pass"`
	School        string `gorm:"column:School"`
	Classes       string `gorm:"column:Classes"`
	Major         string `gorm:"column:Major"`
	Adept         string `gorm:"column:Adept"`
	Vjid          string `gorm:"column:Vjid"`
	Vjpwd         string `gorm:"column:Vjpwd"`
	Email         string `gorm:"column:Email"`
	CodeForceUser string `gorm:"column:CodeForceUser"`
	// HeadUrl string `gorm:"column:HeadUrl"`
}

func (u User) TableName() string {
	return "User"
}

type Permission struct {
	UID          string `gorm:"column:UID"`
	SuperAdmin   string `gorm:"column:SuperAdmin"`
	ProblemAdmin string `gorm:"column:ProblemAdmin"`
	ListAdmin    string `gorm:"column:ListAdmin"`
	SourceAdmin  string `gorm:"column:SourceAdmin"`
	ContestAdmin string `gorm:"column:ContestAdmin"`
}

func (p Permission) TableName() string {
	return "Permission"
}

type Problem struct {
	PID          int    `gorm:"column:PID" json:"PID"`
	Title        string `gorm:"column:Title" json:"Title"`
	Description  string `gorm:"column:Description" json:"Description"`
	Input        string `gorm:"column:Input" json:"Input"`
	Output       string `gorm:"column:Output" json:"Output"`
	SampleInput  string `gorm:"column:SampleInput" json:"SampleInput"`
	SampleOutput string `gorm:"column:SampleOutput" json:"SampleOutput"`
	LimitTime    int64  `gorm:"column:LimitTime" json:"LimitTime"`
	LimitMemory  int64  `gorm:"column:LimitMemory" json:"LimitMemory"`
	Hit          string `gorm:"column:Hit" json:"Hit"`
	Label        string `gorm:"column:Label" json:"Label"`
	Origin       int64  `gorm:"column:Origin" json:"Origin"`
	OriginPID    string `gorm:"column:OriginPID" json:"OriginPID"`
	ContentType  int64  `gorm:"column:ContentType" json:"ContentType"`
}

type List struct {
	LID       int64  `gorm:"column:LID"`
	UID       string `gorm:"column:UID"`
	Title     string `gorm:"column:Title"`
	StartTime int64  `gorm:"column:StartTime"`
}

type ListProblem struct {
	LID   int64  `gorm:"column:LID"`
	PID   int    `gorm:"column:PID"`
	Title string `gorm:"column:Title"`
}

type ListUser struct {
	LID       int64 `gorm:"column:lid"`
	UID       int   `gorm:"column:uid"`
	SubmitNum int   `gorm:"column:SubmitNum"`
	ACNum     int   `gorm:"column:AcNum"`
}

type Contest struct {
	CID         int64  `gorm:"column:CID"`
	UID         string `gorm:"column:UID"`
	Title       string `gorm:"column:Title"`
	Description string `gorm:"column:Description"`
	Begin_time  int64  `gorm:"column:BeginTime"`
	End_time    int64  `gorm:"column:EndTime"`
	Ctype       int    `gorm:"column:Type"`
	Ispublic    int    `gorm:"column:IsPublic"`
	Problems    string `gorm:"column:Problems"`
	Pass        string `gorm:"column:Pass"`
}

func (p Contest) TableName() string {
	return "Contest"
}

type ConPro struct {
	CID        int64  `gorm:"column:CID"`
	PID        int    `gorm:"column:PID"`
	Ptitle     string `gorm:"column:Title"`
	Submit_num int    `gorm:"column:SubmitNum"`
	Ac_num     int    `gorm:"column:ACNum"`
}

func (p ConPro) TableName() string {
	return "ConPro"
}

type Submit struct {
	SID           int                `gorm:"column:SID"`
	PID           int                `gorm:"column:PID"`
	UID           string             `gorm:"column:UID"`
	CID           int                `gorm:"column:CID"`
	Judgeid       int                `gorm:"column:Judgeid"`
	Source        string             `gorm:"column:Source"`
	Lang          constanct.LANG     `gorm:"column:Lang"`
	Result        constanct.OJResult `gorm:"column:Result"`
	Usetime       int64              `gorm:"column:UseTime"`
	UseMemory     int64              `gorm:"column:UseMemory"`
	SubmitTime    int64              `gorm:"column:SubmitTime"`
	IsOriginJudge bool               `gorm:"column:IsOriginJudge"`
	OriginPID     string             `gorm:"column:OriginPID"`
	OJPlatform    int64              `gorm:"column:OJPlatform"`
}

func (p Submit) TableName() string {
	return "Submit"
}

type CeInfo struct {
	SID  int64
	Info string
}

func (p CeInfo) TableName() string {
	return "CEINFO"
}
