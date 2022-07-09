package request

type User struct {
	Uid     string `form:"uid" binding:"required"`
	Uname   string `form:"uname" binding:"required"`
	Pass    string `form:"pass"`
	School  string `form:"school"`
	Classes string `form:"classes"`
	Major   string `form:"major"`
	Vjid    string `form:"vjid"`
	Vjpwd   string `form:"vjpwd"`
	Email   string `form:"email"`
}
