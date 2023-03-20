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
	HeadURL       string `gorm:"column:HeadUrl"`
	Rating        int    `gorm:"column:Rating"`
	LoginIP       string `gorm:"column:LoginIP"`
	RegisterTime  int64  `gorm:"column:RegisterTime"`
	Submited      int64  `gorm:"column:Submited"`
	Solved        uint32 `gorm:"column:Solved"`
	Defaulted     string `gorm:"column:Defaulted"`
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
	PID          string                `gorm:"column:PID" json:"PID" xml:"PID"`
	PType        constanct.ProblemType `gorm:"column:PType" json:"PType" xml:"PType"`
	Title        string                `gorm:"column:Title" json:"Title" xml:"Title"`
	Description  string                `gorm:"column:Description" json:"Description" xml:"Description"`
	Input        string                `gorm:"column:Input" json:"Input" xml:"Input"`
	Output       string                `gorm:"column:Output" json:"Output" xml:"Output"`
	SampleInput  string                `gorm:"column:SampleInput" json:"SampleInput" xml:"SampleInput"`
	SampleOutput string                `gorm:"column:SampleOutput" json:"SampleOutput" xml:"SampleOutput"`
	LimitTime    int64                 `gorm:"column:LimitTime" json:"LimitTime" xml:"LimitTime"`
	LimitMemory  int64                 `gorm:"column:LimitMemory" json:"LimitMemory" xml:"LimitMemory"`
	Hit          string                `gorm:"column:Hit" json:"Hit" xml:"Hit"`
	Label        string                `gorm:"column:Label" json:"Label" xml:"Label"`
	Origin       int64                 `gorm:"column:Origin" json:"Origin" xml:"Origin"`
	OriginPID    string                `gorm:"column:OriginPID" json:"OriginPID" xml:"OriginPID"`
	ContentType  int64                 `gorm:"column:ContentType" json:"ContentType" xml:"ContentType"`
	/*1:可见  -1:不可见*/
	Visible  int    `gorm:"column:Visible" json:"Visible" xml:"Visible"`
	Accepted uint32 `gorm:"column:Accepted" json:"Accepted"`
	Submited uint32 `gorm:"column:Submited" json:"Submited" xml:"Submited"`
	SpjJudge string `gorm:"column:SpjJudge" json:"SpjJudge" xml:"SpjJudge"`
	Source   string `gorm:"column:Source" json:"Source" xml:"Source"`
}

type List struct {
	LID         int64  `gorm:"column:LID"`
	UID         string `gorm:"column:UID"`
	Description string `gorm:"column:Description"`
	Title       string `gorm:"column:Title"`
	StartTime   int64  `gorm:"column:StartTime"`
	Problems    string `gorm:"column:Problems"`
}

func (p List) TableName() string {
	return "List"
}

type ListProblem struct {
	LID      int64  `gorm:"column:LID"`
	PID      string `gorm:"column:PID"`
	Title    string `gorm:"column:Title"`
	Submited uint32 `gorm:"column:Submited"`
	Solved   uint32 `gorm:"column:Solved"`
}

func (p ListProblem) TableName() string {
	return "ListProblem"
}

type ListUser struct {
	LID      int64  `gorm:"column:lid"`
	UID      string `gorm:"column:uid"`
	Submited int    `gorm:"column:Submited"`
	Solved   int    `gorm:"column:Solved"`
}

func (p ListUser) TableName() string {
	return "ListUser"
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
	LangMask    string `gorm:"column:LangMask"`
	Defaulted   string `gorm:"column:Defaulted"`
}

func (p Contest) TableName() string {
	return "Contest"
}

type ConPro struct {
	CID      int64  `gorm:"column:CID"`
	PID      string `gorm:"column:PID"`
	Ptitle   string `gorm:"column:Title"`
	Submited int    `gorm:"column:Submited"`
	Solved   int    `gorm:"column:Solved"`
}

func (p ConPro) TableName() string {
	return "ConPro"
}

type Submit struct {
	SID           int64              `gorm:"column:SID" json:"SID"`
	PID           string             `gorm:"column:PID" json:"PID"`
	UID           string             `gorm:"column:UID" json:"UID"`
	CID           int64              `gorm:"column:CID" json:"CID"`
	JudgeID       int64              `gorm:"column:JudgeID" json:"JudgeID"`
	Source        string             `gorm:"column:Source" json:"Source"`
	Lang          constanct.LANG     `gorm:"column:Lang" json:"Lang"`
	Result        constanct.OJResult `gorm:"column:ResultACM" json:"ResultACM"`
	PassSample    uint32             `gorm:"column:PassSample" json:"PassSample"`
	Sim           uint8              `gorm:"column:Sim" json:"Sim"`
	Usetime       int64              `gorm:"column:UseTime" json:"Usetime"`
	UseMemory     int64              `gorm:"column:UseMemory" json:"UseMemory"`
	SubmitTime    int64              `gorm:"column:SubmitTime" json:"SubmitTime"`
	IsOriginJudge bool               `gorm:"column:IsOriginJudge" json:"IsOriginJudge"`
	OriginPID     string             `gorm:"column:OriginPID" json:"OriginPID"`
	OJPlatform    int64              `gorm:"column:OJPlatform" json:"OJPlatform"`
}

func (p Submit) TableName() string {
	return "Submit"
}

type CeInfo struct {
	SID  int64  `gorm:"column:SID"`
	Info string `gorm:"column:Info"`
}

func (p CeInfo) TableName() string {
	return "CEINFO"
}
