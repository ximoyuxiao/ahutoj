package request

type User struct {
	Uid     string `json:"uid" binding:"required"`
	Uname   string `json:"uname" binding:"required"`
	Pass    string `json:"pass"`
	School  string `json:"school"`
	Classes string `json:"classes"`
	Major   string `json:"major"`
	Vjid    string `json:"vjid"`
	Vjpwd   string `json:"vjpwd"`
	Email   string `json:"email"`
}
type UserInfoReq struct {
	Uid string `json:"uid" binding:"required"`
}

type UserEditReq struct {
	Uname   string `json:"uname" binding:"required"`
	School  string `json:"school"`
	Classes string `json:"classes"`
	Major   string `json:"major"`
	Email   string `json:"email"`
}

type UserEditPassReq struct {
	Fpwd string `json:"fpwd"`
	Spwd string `json:"spwd"`
}

type UserEditVjudgeReq struct {
	Vjid  string `json:"vjid"`
	Vjpwd string `json:"vjpwd"`
}
