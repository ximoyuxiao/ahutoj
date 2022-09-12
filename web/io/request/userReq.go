package request

import "ahutoj/web/dao"

type User struct {
	Uid     string `json:"uid" binding:"required"`
	Uname   string `json:"uname" binding:"required"`
	Pass    string `json:"pass"`
	School  string `json:"school"`
	Classes string `json:"classes"`
	Major   string `json:"major"`
	Adept   string `json:"adept"`
	Vjid    string `json:"vjid"`
	Vjpwd   string `json:"vjpwd"`
	Email   string `json:"email"`
}
type UserInfoReq struct {
	Uid string `json:"uid" binding:"required"`
}
type UserEditReq struct {
	Uname   string `json:"uname"`
	School  string `json:"school"`
	Classes string `json:"classes"`
	Major   string `json:"major"`
	Email   string `json:"email"`
	Adept   string `json:"adept"`
}

type UserEditPassReq struct {
	Pwd    string `json:"pwd" binding:"required"`
	OldPwd string `json:"old_pwd" binding:"required"`
}

type UserEditVjudgeReq struct {
	Vjid  string `json:"vjid"`
	Vjpwd string `json:"vjpwd"`
}

func (u UserEditReq) ToUser(uid string) *dao.User {
	return &dao.User{
		Uid:     uid,
		Uname:   u.Uname,
		School:  u.School,
		Classes: u.Classes,
		Major:   u.Major,
		Email:   u.Email,
		Adept:   u.Adept,
	}
}

func (u UserEditVjudgeReq) ToUser(uid string) *dao.User {
	return &dao.User{
		Uid:   uid,
		Vjid:  u.Vjid,
		Vjpwd: u.Vjpwd,
	}
}
