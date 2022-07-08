package request

type User struct {
	Uid     string `form:"uid"`
	Uname   string `form:"uname"`
	Pass    string `form:"pass"`
	School  string `form:"school"`
	Classes string `form:"classes"`
	Major   string `form:"major"`
	Vjid    string `form:"vjid"`
	Vjpwd   string `form:"vjpwd"`
	Email   string `form:"email"`
}
