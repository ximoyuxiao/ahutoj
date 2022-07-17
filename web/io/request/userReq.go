package request

import (
	"ahutoj/web/dao"
)

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

type UserEditReq struct {
	Uid     string `json:"uid" binding:"required"`
	Uname   string `json:"uname"`
	School  string `json:"school"`
	Classes string `json:"classes"`
	Major   string `json:"major"`
	Email   string `json:"email"`
}

func (u UserEditReq) ToUser() *dao.User {
	return &dao.User{
		Uid:     u.Uid,
		Uname:   u.Uname,
		School:  u.School,
		Classes: u.Classes,
		Major:   u.Major,
		Email:   u.Email,
	}
}

type UserEditPassReq struct {
	Uid     string `json:"uid" binding:"required"`
	Pass    string `json:"pass" binding:"required"`
	OldPass string `json:"old_pass" binding:"required"`
}

// func (u UserEditPassReq) ToUser() *dao.User {
// 	return &dao.User{
// 		Uid:  u.Uid,
// 		Pass: u.Pass,
// 	}
// }

// UserVjudgeBindReq 需要 vj 入参限定
type UserVjudgeBindReq struct {
	Uid   string `json:"uid" binding:"required"`
	Vjid  string `json:"vjid"`
	Vjpwd string `json:"vjpwd"`
}

func (u UserVjudgeBindReq) ToUser() *dao.User {
	return &dao.User{
		Uid:   u.Uid,
		Vjid:  u.Vjid,
		Vjpwd: u.Vjpwd,
	}
}