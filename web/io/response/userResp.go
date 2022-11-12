package response

import (
	"ahutoj/web/dao"
	"ahutoj/web/io/constanct"
)

type LoginResp struct {
	Response
	Token string `json:"Token"`
	Uname string `json:"UserName"`
	Permission
}

type RegisterResp LoginResp

type UserResp struct {
	Response
	UID           string `json:"UID"`
	Uname         string `json:"UserName"`
	School        string `json:"School"`
	Classes       string `json:"Classes"`
	Major         string `json:"Major"`
	Adept         string `json:"Adept"`
	Vjid          string `json:"Vjid"`
	Email         string `json:"Email"`
	CodeForceUser string `json:"CodeForceUser"`
	HeadURL       string `json:"HeadURL"`
}
type UsersItem struct {
	UID      string `json:"UID"`
	Uname    string `json:"UserName"`
	School   string `json:"School"`
	Password string `json:"Password"`
}
type AddUsersResp struct {
	Response
	CreateNumber int         `json:"CreateNumber"`
	Data         []UsersItem `json:"Data"`
}

type UserStatusInfoItem struct {
	PID        string             `json:"PID"`
	SubmitTime int64              `json:"SubmitTime"`
	Result     constanct.OJResult `json:"Result"`
}
type UserStatusInfoResp struct {
	Response
	Data []UserStatusInfoItem `json:"Data"`
}

func CreateUserResp(user *dao.User) UserResp {
	return UserResp{
		Response:      CreateResponse(constanct.SuccessCode),
		UID:           user.UID,
		Uname:         user.Uname,
		School:        user.School,
		Classes:       user.Classes,
		Major:         user.Major,
		Adept:         user.Adept,
		Vjid:          user.Vjid,
		Email:         user.Email,
		CodeForceUser: user.CodeForceUser,
		HeadURL:       user.HeadURL,
	}
}
