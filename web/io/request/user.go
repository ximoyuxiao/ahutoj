package request

import "ahutoj/web/dao"

type User struct {
	UID     string `json:"UID" binding:"required"`
	Uname   string `json:"UserName" binding:"required"`
	Pass    string `json:"Pass"`
	School  string `json:"School"`
	Classes string `json:"Classes"`
	Major   string `json:"Major"`
	Adept   string `json:"Adept"`
	Vjid    string `json:"Vjid"`
	Vjpwd   string `json:"Vjpwd"`
	Email   string `json:"Email"`
}
type UserInfoReq struct {
	UID string `json:"uid" binding:"required"`
}
type UserEditReq struct {
	Uname   string `json:"UserName"`
	School  string `json:"School"`
	Classes string `json:"Classes"`
	Major   string `json:"Major"`
	Email   string `json:"Email"`
	Adept   string `json:"Adept"`
}

type UserEditPassReq struct {
	Pwd    string `json:"Pwd" binding:"required"`
	OldPwd string `json:"OldPwd" binding:"required"`
}

type UserEditVjudgeReq struct {
	Vjid  string `json:"Vjid"`
	Vjpwd string `json:"Vjpwd"`
}
type AddUsersReq struct {
	Number   int     `json:"Number"`
	Prefix   string  `json:"Prefix"`
	School   string  `json:"School"`
	Password *string `json:"Password"`
}

func (u UserEditReq) ToUser(uid string) *dao.User {
	return &dao.User{
		UID:     uid,
		Uname:   u.Uname,
		School:  u.School,
		Classes: u.Classes,
		Major:   u.Major,
		Email:   u.Email,
		Adept:   u.Adept,
	}
}

func (u UserEditVjudgeReq) ToUser(UID string) *dao.User {
	return &dao.User{
		UID:   UID,
		Vjid:  u.Vjid,
		Vjpwd: u.Vjpwd,
	}
}
